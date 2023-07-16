package main

import (
	"fmt"
	"net/url"

	"golang.org/x/net/idna"
)

func main() {
	s := "https://日本語.jp:8443"
	r1, _ := idna.ToASCII(s)
	r2, _ := idna.ToUnicode(r1)

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(url.QueryEscape(s))
}

// xn--https://-5y4qg6h355l.jp:8443
// https://日本語.jp:8443
// https%3A%2F%2F%E6%97%A5%E6%9C%AC%E8%AA%9E.jp%3A8443
