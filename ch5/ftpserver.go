package main

import (
	"log"
	"net"
	"os"
	"strings"
)

const (
	DIR = "DIR"
	CD  = "CD"
	PWD = "PWD"
)

func main() {
	service := "0.0.0.0:1200"
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
	defer conn.Close()
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		s := strings.Split(string(buf[0:n]), " ")
		log.Println(s)

		// декодирование запроса
		switch s[0] {
		case CD:
			chdir(conn, s[1])
		case DIR:
			dirList(conn)
		case PWD:
			pwd(conn)
		default:
			log.Println("Unknown command", s)
		}
	}
}

func chdir(conn net.Conn, s string) {
	if os.Chdir(s) == nil {
		conn.Write([]byte("OK"))
	} else {
		conn.Write([]byte("ERROR"))
	}
}

func pwd(conn net.Conn) {
	s, err := os.Getwd()
	if err != nil {
		conn.Write([]byte(""))
		return
	}
	conn.Write([]byte(s))
}

func dirList(conn net.Conn) {
	// Отправить перенос строки при завершении
	defer conn.Write([]byte("\r\n"))
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	names, err := dir.Readdirnames(-1)
	if err != nil {
		return
	}
	for _, nm := range names {
		conn.Write([]byte(nm + "\r\n"))
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error", err.Error())
	}
}
