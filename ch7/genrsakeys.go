package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := rand.Reader
	bitSize := 2048
	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)
	fmt.Printf("Private key primes: \n[0]:%s\n[1]:%s\n",
		key.Primes[0].String(), key.Primes[1].String())
	fmt.Println("Private key exponent:\n", key.D.String())
	publicKey := key.PublicKey
	fmt.Println("Public key modules:\n", publicKey.N.String())
	fmt.Println("Public key exponent:\n", publicKey.E)
	saveGobKey("ch7/private.key", key)
	saveGobKey("ch7/public.key", publicKey)
	savePEMKey("ch7/private.pem", key)
}

func saveGobKey(fileName string, key any) {
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
}

func savePEMKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	checkError(err)
	var privateKey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	pem.Encode(outFile, privateKey)
	outFile.Close()
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error", err.Error())
	}
}
