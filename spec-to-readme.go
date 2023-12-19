// spec-to-readme converts a JSON part spec into the equivalent documentation
// found in support-docs/README.md. Some manual additions are expected to be
// made after initial generation, but this handles all of the tedious
// copy-pasting between files.
//
// Simple usage:
//
//    cat path/to/partspec.json | go run spec-to-readme.go --section-num="3.123"
//
// Complex usage (bash):
/*
   CUR=10
   for A in example1.json example2.json example3.json; do
      cat part-spec/$A | go run spec-to-readme.go --section-num="3.$CUR"
      ((CUR++))
   done  >> support-docs/README.md
*/
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
)

var sectionNum = flag.String("section-num", "3.XX", "Section title prefix")

type topLevel struct {
	// TODO: double-check that $id fields match the filename

	// If this top-level object type can be satisfied by other object types
	OneOf []componentOption

	// If this top-level object type just has specific properties of its own.
	componentOption
}

type componentOption struct {
	Title      string
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
			log.Fatal("Cannot set both Ref and Type")
		}
		return p.Ref
	}
	if p.Type == "" {
		return "Must set either Ref or Type"
	}
	switch p.Type {
	case "array":
		if p.Items == nil {
			log.Fatal("Must set 'items' field when setting type=field")
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
	flag.Parse()
	decodeIn := json.NewDecoder(os.Stdin)

	for {
		var top topLevel
		if err := decodeIn.Decode(&top); err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
			log.Fatal(err)
		}

		if len(top.OneOf) > 0 {
			processOneOf(top)
		} else {
			opt := top.componentOption
			if len(opt.Properties) == 0 {
				panic(fmt.Sprintf("invalid top-level object %q: no oneOf or properties field found", top.Title))
			}
			processComponentOption(*sectionNum, opt)
		}
	}

	panic("unreachable")
}

func processOneOf(top topLevel) {
	fmt.Printf("### %s\t %s\n\n", *sectionNum, titleCase(top.Title))

	for i, opt := range top.OneOf {
		idx := i + 1
		prefix := fmt.Sprintf(" %s.%d", *sectionNum, idx)
		processComponentOption(prefix, opt)
	}
}

func processComponentOption(titlePrefix string, opt componentOption) {
	requiredSet := make(map[string]bool)
	for _, r := range opt.Required {
		requiredSet[r] = true
	}
	propList, propMap := processProperties(opt.Properties)
	fmt.Printf("#### %s\t %s\n\n", titlePrefix, titleCase(opt.Title))
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

func processProperties(raw json.RawMessage) ([]string, map[string]property) {
	m := make(map[string]property)
	if err := json.Unmarshal(raw, &m); err != nil {
		panic(err)
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
	"wlan", "WLAN",
}

func titleCase(sentence string) string {
	var out []string
	for _, word := range strings.Split(sentence, " ") {

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
				panic("one letter word? " + sentence)
			}
			upperWord = string(unicode.ToTitle(firstChar)) + word[secondCharStart:]
		}

		out = append(out, upperWord)
	}

	return strings.Join(out, " ")
}
