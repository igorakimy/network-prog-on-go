package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage:", os.Args[0], "host:port")
	}
	service := os.Args[1]
	// Подключиться по защищенному каналу
	conn, err := tls.Dial("tcp", service, nil)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := io.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	conn.Close()
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error", err.Error())
	}
}
