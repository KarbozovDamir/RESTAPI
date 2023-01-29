package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	jsFile, err := os.Open("users.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsFile.Close()
	fmt.Println("Successfully loaded")

	// read content of jsFile as bytes
	byteValue, err := ioutil.ReadAll(jsFile)
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]interface{}
	// mv from byteValue to users
	json.Unmarshal(byteValue, &result)
	fmt.Println(result["users"])

}
