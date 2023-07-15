package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

func main() {
	// Загружаем сертификат, чтобы мы могли получить
	// доступ к встроенному открытому ключу
	certCerFile, err := os.Open("ch7/jan.newmarch.name.cer")
	checkError(err)
	// Больше, чем размер файла
	derBytes := make([]byte, 1000)
	count, err := certCerFile.Read(derBytes)
	checkError(err)
	certCerFile.Close()
	// Обрезаем байты до фактической длины при вызове
	cert, err := x509.ParseCertificate(derBytes[0:count])
	checkError(err)
	fmt.Printf("Name %s\n", cert.Subject.CommonName)
	fmt.Printf("Not before %s\n", cert.NotBefore.String())
	fmt.Printf("Not after %s\n", cert.NotAfter.String())

	// Загружаем невстроенный публичный ключ, который должен
	// быть таки же, как встроенный ключ выше
	pub, err := os.Open("ch7/public.key")
	checkError(err)
	dec := gob.NewDecoder(pub)
	publicKey := new(rsa.PrivateKey)
	err = dec.Decode(publicKey)
	checkError(err)
	pub.Close()

	// genx509cert.go создал открытый ключ и сертификат
	// сертификаты также содержат открытый ключ
	// мы сравниваем публичный ключ и поля встроенного открытого ключа
	if cert.PublicKey.(*rsa.PublicKey).N.Cmp(publicKey.N) == 0 {
		if publicKey.E == cert.PublicKey.(*rsa.PublicKey).E {
			fmt.Println("Same public key")
			return
		}
	}
	fmt.Println("Different public key")
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error", err.Error())
	}
}
