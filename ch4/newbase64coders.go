package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
)

var (
	enc   = base64.StdEncoding.WithPadding('|')
	input = []byte("∞a∞\x02ab")
	w     bytes.Buffer
)

func restoreViaDecoder() {
	var buf = bytes.NewBuffer(w.Bytes())
	var ior = base64.NewDecoder(enc, buf)
	l := len(input)

	if l > 3 && l%3 != 0 {
		l += 2
	}

	restored := make([]byte, l)
	ior.Read(restored)
	fmt.Printf("%11s: %s %v\n", "viaDecoder", string(restored), restored)
}

func restoreViaEncoding() {
	var dst = make([]byte, len(input))
	enc.Decode(dst, w.Bytes())
	fmt.Printf("%11s: %s %v\n", "viaEncoding", string(dst), dst)
}

func main() {
	fmt.Printf("%11s: %s %v\n", "input", string(input), input)

	var wc = base64.NewEncoder(enc, &w)

	wc.Write(input)
	wc.Close()

	fmt.Printf("%11s: %s %v\n", "encoded", string(w.Bytes()), w.Bytes())

	restoreViaDecoder()
	restoreViaEncoding()
}
