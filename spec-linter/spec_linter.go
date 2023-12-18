package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"maps"
	"os"
	"path/filepath"
	"strings"
)

var (
	digitalDatasheetDir = flag.String("datasheet-dir", "", "Locational of digital-datasheets repo on local machine")
)

func main() {
	flag.Parse()

	fileList := []string{}
	err := filepath.Walk(*digitalDatasheetDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip .git and part-spec-proto directories
		if info.IsDir() && (strings.Contains(info.Name(), ".git") || strings.Contains(info.Name(), "part-spec-proto")) {
			return filepath.SkipDir
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
			fileList = append(fileList, path)
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, file := range fileList {
		fmt.Printf("\n\n%s\n", file)
		b, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}
		var anyJson map[string]interface{}
		err = json.Unmarshal(b, &anyJson)
		if err != nil {
			panic(err)
		}

		// struct fields in the digital datasheet schema
		keyList := maps.Keys(anyJson)

		// each file should contain "$schema"
		v, ok := anyJson["$schema"]
		if !ok {
			panic(fmt.Errorf("expecting $schema key in %v", keyList))
		}

		// $title should start with "specification"
		v, ok = anyJson["title"]
		if !ok {
			panic(fmt.Errorf("expecting title key in %v", keyList))
		}
		title := fmt.Sprintf("%v", v)
		if strings.HasPrefix(title, "Specification") {
			anyJson["title"] = strings.ToLower(title)
			fmt.Printf("fixed title: %s\n", anyJson["title"])
		} else if !strings.HasPrefix(title, "specification") {
			anyJson["title"] = fmt.Sprintf("specification of %s", title)
			fmt.Printf("fixed title: %s\n", anyJson["title"])
		}

		// $id should have this format: "https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/clock/oscillator.json"
		v, ok = anyJson["$id"]
		if !ok {
			panic(fmt.Errorf("expecting $id key in %v", keyList))
		}
		idString := fmt.Sprintf("%v", v)
		prefixPath := "https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/"
		correctPrefix := strings.HasPrefix(idString, prefixPath)
		filePieces := strings.Split(file, "/")
		suffixPath := filepath.Join(filePieces[len(filePieces)-2], filePieces[len(filePieces)-1])
		correctPath := strings.HasSuffix(idString, suffixPath)
		if !correctPrefix || !correctPath {
			anyJson["$id"] = filepath.Join(prefixPath, suffixPath)
			fmt.Printf("fixed $id: %s\n", anyJson["$id"])
		}

		// For each property key:
		//		1. Property name should be lowercase
		//		2. Description should be lowercase
		//		3. Description should not end in a period
		//		// TODO: this has currently been fixed by hand, but we should add this to the script
		//		4. Ref should go to an actual place

		// Drill into the structure to get to the component properties
		objectName := strings.TrimSuffix(filepath.Base(file), ".json")
		v, ok = anyJson[objectName]
		if !ok {
			fmt.Printf("expecting %s key in %v\n", objectName, keyList)
			continue
		}
		vMap, err := castToMap(v)
		if err != nil {
			fmt.Println(err)
			continue
		}

		v, ok = vMap["properties"]
		if !ok {
			fmt.Printf("expecting properties key in %v\n", maps.Keys(vMap))
		}
		vMap, err = castToMap(v)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for k, v := range vMap {
			// make sure the property names start with a lowercase letter
			if checkForCapitalWithSkips(k) {
				newK := strings.ToLower(string(k[0])) + k[1:]
				vMap[newK] = v
				delete(vMap, k)
				fmt.Printf("fixed property name: %s\n", newK)
			}
		}

		for _, v := range vMap {
			// For each property name, the description field should start with a lowercase letter and NOT end in a period
			propMap, err := castToMap(v)
			if err != nil {
				fmt.Println(err)
				continue
			}
			desc, ok := propMap["description"]
			if !ok {
				fmt.Printf("expecting description key in %v\n", maps.Keys(propMap))
			}
			descStr := fmt.Sprintf("%v", desc)
			if checkForCapitalWithSkips(descStr) {
				newDesc := strings.ToLower(string(descStr[0])) + descStr[1:]
				propMap["description"] = newDesc
				// needed to modify below
				descStr = newDesc
				fmt.Printf("fixed capital: %s\n", newDesc)
			}

			if string(descStr[len(descStr)-1]) == "." {
				newDesc := descStr[0 : len(descStr)-1]
				propMap["description"] = newDesc
				fmt.Printf("fixed period: %s\n", newDesc)
			}
		}

		newB, err := json.Marshal(anyJson)
		if err != nil {
			panic(err)
		}
		err = os.WriteFile(file, newB, 0775)
		if err != nil {
			panic(err)
		}

	}

}

func castToMap(v any) (map[string]interface{}, error) {
	vMap, ok := v.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("cannot type cast %+v to a map", v)
	}
	return vMap, nil
}

func checkForCapitalWithSkips(s string) bool {
	skipMap := map[string]bool{"LED": true, "SNR": true, "URI": true, "DC ": true}
	if len(s) < 3 {
		return false
	}
	if _, ok := skipMap[s[0:3]]; ok {
		return false
	}
	return strings.ToUpper(string(s[0])) == string(s[0])
}
