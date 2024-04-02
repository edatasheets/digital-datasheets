// spec-to-readme converts a JSON part spec into the equivalent documentation
// found in support-docs/README.md. Some manual additions are expected to be
// made after initial generation, but this handles all of the tedious
// copy-pasting between files.
//
// This script expects two arguments: a directory with json schema files and the section number
//
// Simple usage:
//
//    cat path/to/partspec.json | go run spec-to-readme.go  part-spec/passive "3.123"
//
// Complex usage (bash):
/*
   CUR=10
   for A in passives/ semiconductors/ ic_mic/; do
      go run spec-to-readme.go part-spec/$A "3.$CUR"
      ((CUR++))
   done  >> support-docs/README.md
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
)

type topLevel struct {
	ID    string `json:"$id"`
	Title string `json:"title"`

	// If this top-level object type just has specific properties of its own.
	X map[string]json.RawMessage

	// The top-level object can have defs of its own.
	Defs map[string]componentOption `json:"$defs"`
}

type componentOption struct {
	Type       string
	Required   []string
	Properties json.RawMessage
}

type property struct {
	Description string
	// For a non-compound type
	Ref string `json:"$ref"`
	// For compount types
	Type  string
	Items *property
}

func (p property) TypeString() string {
	if p.Ref != "" {
		if p.Type != "" {
			log.Fatalf("Bad property %q: cannot set both Ref (%q) and Type (%q)", p.Description, p.Ref, p.Type)
		}
		return p.Ref
	}
	if p.Type == "" {
		return "Must set either Ref or Type"
	}
	switch p.Type {
	case "array":
		if p.Items == nil {
			log.Fatal("Must set 'items' field when setting type=array")
		}
		return "array of " + p.Items.TypeString()
	case "number":
		return "Number"
	case "string":
		return "String"
	case "boolean":
		return "Boolean"
	default:
		log.Fatalf("Unknown property Type %q", p.Type)
	}
	panic("unreachable")

}

func main() {
	// Make a list of all the json files in this part type directory
	rootDir := os.Args[1]
	fileList := []string{}
	err := filepath.Walk(rootDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
			fileList = append(fileList, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	// The section title is the same as the directory name
	sectionNum := os.Args[2]
	title := strings.TrimSuffix(strings.TrimPrefix(rootDir, "part-spec/"), "/")
	fmt.Printf("### %s\t %s\n\n", sectionNum, titleCase(title))

	i := 1
	for _, f := range fileList {
		b, err := os.ReadFile(f)
		if err != nil {
			panic(err)
		}
		// Unmarshal known fields
		var top topLevel
		if err := json.Unmarshal(b, &top); err != nil {
			panic(err)
		}

		// Unmarshal extra fields (really things like speaker_amplifier, bjt, etc)
		if err := json.Unmarshal(b, &top.X); err != nil {
			panic(err)
		}
		deleteList := []string{"$id", "$schema", "title", "$defs"}
		for _, d := range deleteList {
			delete(top.X, d)
		}
		if len(top.X) > 1 {
			panic(fmt.Sprintf("there should only be one property object per file"))
		}
		// Unmarshal the extra field (bjt, mosfet, etc) into a componentOption
		var opt componentOption
		for _, v := range top.X {
			json.Unmarshal(v, &opt)
		}

		// Process component values and $defs
		if len(opt.Properties) > 0 {
			fmt.Printf("")
			prefix := fmt.Sprintf(" %s.%d", sectionNum, i)
			_, shortName := path.Split(top.ID)
			processComponentOption(prefix, top.Title, source{shortName: shortName, uri: top.ID}, opt)
			i += 1
		} else if len(top.Defs) == 0 {
			fmt.Fprintf(os.Stderr, "invalid top-level object %q: no oneOf, properties, or definitions field found\n", top.Title)
		}
		if len(top.Defs) > 0 {
			i = processDefs(" "+sectionNum, i, top.Defs)
			i += 1

		}

	}
}

type source struct {
	shortName string
	uri       string
}

func processComponentOption(titlePrefix, title string, link source, opt componentOption) {
	requiredSet := make(map[string]bool)
	for _, r := range opt.Required {
		requiredSet[r] = true
	}
	propList, propMap := processProperties(opt.Properties)
	if len(propList) == 0 {
		panic(fmt.Sprintf("no properties found for %q", title))
	}
	fmt.Printf("#### %s\t %s\n\n", titlePrefix, titleCase(title))
	if (link != source{}) {
		fmt.Printf("Source: [%v](%v)\n\n", link.shortName, link.uri)
	}
	fmt.Println("|Property|Description|JSON Data Type|Required?|")
	fmt.Println("|:----|:----|:----|:----|")
	for _, name := range propList {
		prop := propMap[name]
		required := " "
		if requiredSet[name] {
			required = "Yes"
		}
		fmt.Printf("|%s|%s|%s|%s|\n", name, prop.Description, prop.TypeString(), required)
	}
	fmt.Println()
}

func processDefs(titlePrefix string, idOffset int, defs map[string]componentOption) int {
	if len(defs) == 0 {
		return idOffset
	}

	var subDefNames []string
	for name := range defs {
		subDefNames = append(subDefNames, name)
	}
	sort.Strings(subDefNames)
	idx := 0
	for i, subDefName := range subDefNames {
		subDef := defs[subDefName]
		idx = i + idOffset
		processComponentOption(fmt.Sprintf("%s.%d", titlePrefix, idx), subDefName, source{}, subDef)
	}
	return idx
}

func processProperties(raw json.RawMessage) ([]string, map[string]property) {
	if len(raw) == 0 {
		return nil, nil
	}
	m := make(map[string]property)
	if err := json.Unmarshal(raw, &m); err != nil {
		panic(fmt.Sprintf("%s: %s", err, raw))
	}

	rawString := string(raw)
	ordering := make(map[string]int, len(m))
	nameList := make([]string, 0, len(m))
	for name := range m {
		nameList = append(nameList, name)

		// HACK: to preserve the ordering of the properties, we have to do some gross stuff.
		lookFor := `"` + name + `": {`
		leftIdx := strings.Index(rawString, lookFor)
		rightIdx := strings.LastIndex(rawString, lookFor)
		if leftIdx != rightIdx {
			panic("Try harder: " + lookFor + " was found multiple times")
		}
		ordering[name] = leftIdx
	}

	sort.Slice(nameList, func(i, j int) bool {
		l, lok := ordering[nameList[i]]
		if !lok {
			panic(fmt.Sprintf("Couldn't find %q in ordering", nameList[i]))
		}
		r, rok := ordering[nameList[j]]
		if !rok {
			panic(fmt.Sprintf("Couldn't find %q in ordering", nameList[j]))
		}
		return l < r
	})
	return nameList, m
}

var upperAcronyms = map[string]string{
	"io":   "IO",
	"ic":   "IC",
	"sd":   "SD",
	"ssd":  "SSD",
	"wwan": "WWAN",
	"wlan": "WLAN",
}

func titleCase(sentence string) string {
	var out []string
	for _, word := range strings.Fields(sentence) {

		upperWord, ok := upperAcronyms[word]
		if !ok {
			var firstChar rune
			var secondCharStart int
			for i, ch := range word {
				if i == 0 {
					firstChar = ch
					continue
				}

				secondCharStart = i
				break
			}
			if secondCharStart == 0 {
				return sentence
				// panic("one letter word? " + sentence)
			}
			upperWord = string(unicode.ToTitle(firstChar)) + word[secondCharStart:]
		}

		out = append(out, upperWord)
	}

	return strings.Join(out, " ")
}
