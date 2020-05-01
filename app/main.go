package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/go-cmp/cmp"
	"github.com/toshick/package-lock-json-diff/app/model"
)

var url1 = "../data/before/package-lock.json"
var url2 = "../data/after/package-lock.json"

func main() {
	fmt.Println("\npackage-lock.jsonの差分を検出しよう\n")

	// structによる差分チェック
	if err := diffJsonStruct(); err != nil {
		log.Fatal(err)
		return
	}

	// interface{}による差分チェック
	if err := diffJsonInterface(); err != nil {
		log.Fatal(err)
		return
	}
}

/*
 * interface{}による差分チェック
 */
func diffJsonInterface() error {
	json1 := getJsonInterface(url1)
	json2 := getJsonInterface(url2)

	fmt.Printf("-------- interface diff --------' \n")

	dep1 := json1["dependencies"]
	dep2 := json2["dependencies"]

	for key := range dep1.(map[string]interface{}) {
		if diff := cmp.Diff(dep1.(map[string]interface{})[key], dep2.(map[string]interface{})[key]); diff != "" {
			fmt.Printf("変更あり '%v' \n", key)
		}
	}
	return nil
}

/*
 * structによる差分チェック
 */
func diffJsonStruct() error {
	json1 := getJsonStruct(url1)
	json2 := getJsonStruct(url2)

	fmt.Printf("-------- struct diff --------' \n")
	for key := range json1.Dependencies {
		// fmt.Printf("キー %v \n", json1.Dependencies[key])
		dep1 := json1.Dependencies[key]
		dep2 := json2.Dependencies[key]

		if diff := cmp.Diff(dep1, dep2); diff != "" {
			// fmt.Printf(" differs: (%v)\n%s", key, diff)
			fmt.Printf("変更あり '%v' \n", key)
		}
	}
	return nil
}

/*
 * getJsonStruct
 */
func getJsonStruct(url string) *model.PackageLockJson {
	jsonstr := getFile(url)
	// fmt.Printf("l %v \n", j)

	pglj := new(model.PackageLockJson)

	if err := json.Unmarshal([]byte(jsonstr), pglj); err != nil {
		log.Fatal(err)
	}
	return pglj
}

/*
 * getJsonInterface
 */
func getJsonInterface(url string) map[string]interface{} {
	jsonstr := getFile(url)
	// fmt.Printf("l %v \n", j)

	pglj := map[string]interface{}{}

	if err := json.Unmarshal([]byte(jsonstr), &pglj); err != nil {
		log.Fatal(err)
	}
	return pglj
}

/*
 * getFile
 */
func getFile(url string) string {
	bytes, err := ioutil.ReadFile(url)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

/*
 * saveFile
 */
func saveFile(data *model.PackageLockJson) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	url := "../data/sample.json"
	if err := ioutil.WriteFile(url, bytes, 0); err != nil {
		return err
	}

	err = os.Chmod(url, 0777)
	if err != nil {
		return err
	}

	return nil
}
