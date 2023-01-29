package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Age  int    `json:"age"`
	Job  Job    `json:"job"`
}

type Job struct {
	Backend  string `json:"backend"`
	Frontend string `json:"frontend"`
}

func main() {
	jsFile, err := os.Open("users.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsFile.Close()
	fmt.Println("Successfully loaded")

	var users Users
	// read content of jsFile as bytes
	byteValue, err := ioutil.ReadAll(jsFile)
	if err != nil {
		log.Fatal(err)
	}
	// mv from byteValue to users
	json.Unmarshal(byteValue, &users)
	for _, u := range users.Users {
		fmt.Println("------")
		PrintUser(&u)
	}
}

func PrintUser(u *User) {
	fmt.Printf("Name: %s\n", u.Name)
	fmt.Printf("Type: %s\n", u.Type)
	fmt.Printf("Age: %d\n", u.Age)
	fmt.Printf("Job: %s\n", u.Job.Backend)
	fmt.Printf("Job: %s\n", u.Job.Frontend)
}
