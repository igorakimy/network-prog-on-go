package main

import (
	"log"
	"net"
	"time"
)

func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		// Записать текущую дату в текущее установленное соединение.
		conn.Write([]byte(daytime))
		// Закрыть текущее соединение.
		conn.Close()
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s\n", err.Error())
	}
}
