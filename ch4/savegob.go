package main

import (
	"encoding/gob"
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
	saveGob("ch4/person.gob", person)
}

func saveGob(fileName string, key any) {
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error", err.Error())
	}
}
