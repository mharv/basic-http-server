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
func SaveDb(content map[string]interface{}) string {
	contentJson, err := json.Marshal(content)

	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile("fakeDb.json", contentJson, 0666)

	if err != nil {
		fmt.Println(err)
	}
	return "save complete\n"
}
func DeleteDb() string {

	err := os.WriteFile("fakeDb.json", []byte{}, 0666)

	if err != nil {
		fmt.Println(err)
	}
	return "delete complete\n"
}
