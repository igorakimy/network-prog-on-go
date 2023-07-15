package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/gob"
	"encoding/pem"
	"log"
	"math/big"
	"os"
	"time"
)

func main() {
	random := rand.Reader
	var key rsa.PrivateKey
	loadKey("ch7/private.key", &key)
	now := time.Now()
	then := now.Add(60 * 60 * 24 * 365 * 1000 * 1000 * 1000)
	// Один год
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName:   "jan.newmarch.name",
			Organization: []string{"Jan Newmarch"},
		},
		NotBefore:    now,
		NotAfter:     then,
		SubjectKeyId: []byte{1, 2, 3, 4},
		KeyUsage: x509.KeyUsageCertSign |
			x509.KeyUsageKeyEncipherment |
			x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		IsCA:                  true,
		DNSNames:              []string{"jan.newmarch.name", "localhost"},
	}
	derBytes, err := x509.CreateCertificate(random, &template,
		&template, &key.PublicKey, &key)
	checkError(err)
	certCerFile, err := os.Create("ch7/jan.newmarch.name.cer")
	checkError(err)
	certCerFile.Write(derBytes)
	certCerFile.Close()
	certCerFile.Close()
	certPEMFile, err := os.Create("ch7/jan.newmarch.name.pem")
	checkError(err)
	pem.Encode(certPEMFile, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derBytes,
	})
	certPEMFile.Close()
	keyPEMFile, err := os.Create("ch7/private.pem")
	checkError(err)
	pem.Encode(keyPEMFile, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(&key),
	})
	keyPEMFile.Close()
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
		log.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}
