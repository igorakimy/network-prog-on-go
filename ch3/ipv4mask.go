package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s dotted-ip-addr\n", os.Args[0])
	}
	dotAddr := os.Args[1]
	addr := net.ParseIP(dotAddr)
	if addr == nil {
		log.Fatalln("nil Invalid address")
	}

	// Получить IP-маску по умолчанию для IPv4-адреса
	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()
	fmt.Println("Address is", addr.String(),
		"\nDefault mask length is", bits,
		"\nLeading ones count is", ones,
		"\nMask is (hex)", mask.String(),
		"\nNetwork is", network.String())

	// Получить IPv4 маску в 4-байтовой форме
	derivedMask := net.IPv4Mask(255, 255, 0, 0)
	fmt.Printf("Network using %s: %s\n", derivedMask, addr.Mask(derivedMask))
}
