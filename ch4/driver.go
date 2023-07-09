package main

import (
	"encoding/asn1"
	"fmt"

	p "badtype"
)

func main() {
	t1 := p.T{F: 1}
	mdata1, err := asn1.Marshal(t1)
	fmt.Println(err)
	var newT1 p.T
	_, err = asn1.Unmarshal(mdata1, &newT1)
	fmt.Println(err)
}
