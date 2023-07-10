package main

import (
	"encoding/asn1"
	"fmt"
	"log"
)

type MyType struct {
	F1 rune
	F2 int
}

type YourType struct {
	F3 rune
}

type TheirType struct {
	F4 byte
}

func main() {
	// Первый рабочий пример
	t1 := MyType{'ロ', 1}
	mdata1, _ := asn1.Marshal(t1)
	t2 := new(YourType)
	_, err := asn1.Unmarshal(mdata1, t2)
	fmt.Printf("Before marshal: %v, after unmarshal: %v\n", t1, t2)
	checkError(err)

	// Синтаксическая ошибка (не удается заполнить все поля)
	y := YourType{'ロ'}
	mdata2, _ := asn1.Marshal(y)
	z := new(MyType)
	_, err = asn1.Unmarshal(mdata2, z)
	fmt.Printf("Before marshal: %v, after unmarshal: %v\n", y, z)
	checkError(err)

	// Структурная ошибка (неправильный тип Go byte != rune)
	t3 := new(TheirType)
	_, err = asn1.Unmarshal(mdata1, t3)
	fmt.Printf("Before marshal: %v, after unmarshal: %v\n", t1, t3)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
