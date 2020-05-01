package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/toshick/package-lock-json-diff/app/model"
)

func main() {
	fmt.Println("Hello, 世界")
	// pkgl := &c{}
	// fmt.Printf("pkgl %v \n", pkgl)

	jsonstr := getFile("../data/after/package-lock.json")
	// fmt.Printf("l %v \n", j)

	pglj := new(model.PackageLockJson)
	err := json.Unmarshal([]byte(jsonstr), pglj)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("pglj %v \n", pglj)
	saveFile(pglj)
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
		panic(err)
	}
	// str := string(bytes)
	// fmt.Println(string)

	if err := ioutil.WriteFile("../data/sample.json", bytes, 0); err != nil {
		log.Fatalln(err)
	}

	return nil
}
