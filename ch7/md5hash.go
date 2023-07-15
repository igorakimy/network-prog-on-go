package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	hash := md5.New()
	bytes := []byte("hello\n")
	// Добавляем данные в работающий хэш
	hash.Write(bytes)
	// Получаем хэшированные данные
	hashValue := hash.Sum(nil)
	// Количество байт возвращаемых суммой
	hashSize := hash.Size()
	// Для каждых 4 байт hashValue
	// мы вставляем в байт val со сдвигом
	// val[first_byte] = hashValue[n] после сдвига 24
	// вторая и третья позиции байта после 16 и 8
	// val[fourth_byte] = hashValue[n+3]
	// в итоге у нас есть значение uint32, которое мы печатаем
	for n := 0; n < hashSize; n += 4 {
		var val uint32
		val = uint32(hashValue[n])<<24 +
			uint32(hashValue[n+1])<<16 +
			uint32(hashValue[n+2])<<8 +
			uint32(hashValue[n+3])
		fmt.Printf("%x ", val)
	}
	fmt.Println()
}
