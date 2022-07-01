package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ConnectDb() map[string]interface{} {
	jsonFile, err := os.Open("fakeDb.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened faeDb.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return result
}
