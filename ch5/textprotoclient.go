package main

import (
	"fmt"
	"log"
	"net/textproto"
)

func main() {
	conn, err := textproto.Dial("unix", "/tmp/fakewebserver")
	checkError(err)
	defer conn.Close()
	fmt.Println("Sending request to retrieve /mypage")
	id, err := conn.Cmd("GET /mypage")
	checkError(err)
	conn.StartResponse(id)
	defer conn.EndResponse(id)
	// Фейковая отправка 200 через nc или ваш собственный сервер
	code, stringResult, err := conn.ReadCodeLine(200)
	checkError(err)
	fmt.Println(code, "\n", stringResult, "\n", err)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error", err.Error())
	}
}
