package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Professor struct {
	Name      string `json:"name"`
	ScienceID int    `json:"science_id"`
	IsWorking bool   `json:"is_orking"`
	Univer    Univer `json:"univer"`
}

type Univer struct {
	Name string `json:"name"`
	City string `json:"city"`
}

func main() {
	prof1 := Professor{
		Name:      "Damir",
		ScienceID: 12,
		IsWorking: true,
		Univer: Univer{
			Name: "MIT",
			City: "Ata",
		},
	}
	// Change prof1 to bytes(JSON)
	byteArr, err := json.MarshalIndent(prof1, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(byteArr))
	os.WriteFile("output.json", byteArr, 0666) //rwrwrw
	if err != nil {
		log.Fatal(err)
	}
}
