package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	// Закрыть соединение при выходе
	defer conn.Close()
	var buf [512]byte
	for {
		// Чтение до 512 байт
		n, err  := conn.Read(buf[0:])
		checkError(err)
		fmt.Println(string(buf[0:]))
		// Записать n-прочитанных байт
		_, err = conn.Write(buf[0:n])
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s\n", err.Error())
	}
}