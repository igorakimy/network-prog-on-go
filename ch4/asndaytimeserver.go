package main

import (
	"encoding/asn1"
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
		daytime := time.Now()
		mdata, _ := asn1.Marshal(daytime)
		conn.Write(mdata)
		conn.Close()
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %v\n", err.Error())
	}
}
