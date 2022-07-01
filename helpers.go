package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func GetDb() map[string]interface{} {
	jsonFile, err := os.Open("fakeDb.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened fakeDb.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return result
}
