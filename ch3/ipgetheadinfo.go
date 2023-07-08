package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Using: %s host:port", os.Args[0])
	}
	service := os.Args[1]
	// Подключиться к адресу в указанной сети
	conn, err := net.Dial("tcp", service)
	checkError(err)
	// Отправить байты на уснановленное соединение
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := readFully(conn)
	checkError(err)
	fmt.Println(string(result))

}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	// Создать новый буфер для записи в него данных
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		// Считать байты в переменную
		n, err := conn.Read(buf[0:])
		// Записать байты из переменной в буфер
		result.Write(buf[0:n])
		if err != nil {
			// Если данные были полностью прочитаны
			if err == io.EOF {
				// Завершишь выполнение цикла
				break
			}
			return nil, err
		}
	}
	// Прочитать байты из буфера и вернуть их
	return result.Bytes(), nil
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s", err.Error())
	}
}
