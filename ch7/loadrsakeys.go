package main

import (
	"crypto/rsa"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

func main() {
	var key rsa.PrivateKey
	loadKey("ch7/private.key", &key)
	fmt.Printf("Private kee primes:\n[0]:%s\n[1]:%s\n",
		key.Primes[0].String(), key.Primes[1].String())
	fmt.Println("Private key exponent:\n", key.D.String())
	var publicKey rsa.PublicKey
	loadKey("ch7/public.key", &publicKey)
	fmt.Println("Public key modules:\n", publicKey.N.String())
	fmt.Println("Public key exponent:\n", publicKey.E)
}

func loadKey(fileName string, key any) {
	inFile, err := os.Open(fileName)
	checkError(err)
	decoder := gob.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkError(err)
	inFile.Close()
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error", err.Error())
	}
}
