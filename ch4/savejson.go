package main

import (
	"encoding/json"
	"log"
	"os"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}

func main() {
	person := Person{
		Name: Name{Family: "Newmarch", Personal: "Jan"},
		Email: []Email{
			{Kind: "home", Address: "jan@newmarch.name"},
			{Kind: "work", Address: "j.newmarch@boxhill.edu.au"},
		},
	}
	saveJSON("ch4/person.json", person)
}

func saveJSON(fileName string, key any) {
	data, err := json.Marshal(key)
	checkError(err)
	err = os.WriteFile(fileName, data, 0600)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error", err.Error())
	}
}
