package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"
)

func main() {
	rootPEM, err := os.ReadFile("ch7/jan.newmarch.name.pem")
	checkError(err)
	// Сначала создадим набор коневых сертификатов. Для этого примера мы
	// имеем только один. Это также можно опустить, чтобы использовать
	// корневой набор по умолчанию для текущей операционной системы
	roots := x509.NewCertPool()
	if ok := roots.AppendCertsFromPEM(rootPEM); !ok {
		panic("failed to parse root certificate")
	}
	conn, err := tls.Dial("tcp", "localhost:1200", &tls.Config{
		RootCAs: roots,
	})
	if err != nil {
		panic("failed to connect: " + err.Error())
	}

	// Запись и чтение несколько раз подряд
	for n := 0; n < 10; n++ {
		fmt.Println("Writing...")
		conn.Write([]byte("Hello " + string(rune(n+48))))
		var buf [512]byte
		n, _ := conn.Read(buf[0:])
		fmt.Println(string(buf[0:n]))
	}

	conn.Close()
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error", err.Error())
	}
}
