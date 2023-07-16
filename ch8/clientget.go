package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage:", os.Args[0], "http://host:port/page")
	}
	URL, err := url.Parse(os.Args[1])
	checkError(err)
	client := &http.Client{}
	request, err := http.NewRequest("HEAD", URL.String(), nil)
	checkError(err)
	response, err := client.Do(request)
	checkError(err)
	if response.StatusCode != http.StatusOK {
		log.Fatalln(response.Status)
	}
	fmt.Println("The response header is")
	b, _ := httputil.DumpResponse(response, false)
	fmt.Println(string(b))
	chSet := getCharset(response)
	if chSet != "utf-8" {
		log.Fatalln("Cannot handle", chSet)
	}
	var buf [512]byte
	reader := response.Body
	fmt.Println("got body")
	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			if err == io.EOF {
				fmt.Print(string(buf[0:]))
				break
			}
			checkError(err)
		}
		fmt.Print(string(buf[0:n]))
	}
}

func getCharset(response *http.Response) string {
	contentType := response.Header.Get("Content-Type")
	if contentType == "" {
		return "utf-8"
	}
	idx := strings.Index(contentType, "charset=")
	if idx == -1 {
		return "utf-8"
	}
	// мы нашли кодировку, теперь вырезаем ее
	chSet := strings.Trim(contentType[idx+8:], " ")
	return strings.ToLower(chSet)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error", err.Error())
	}
}
