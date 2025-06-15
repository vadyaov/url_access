package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/vadyaov/url_access/urlparser"
)

var fileFlag string

func init() {
	const (
		noFileFlag = ""
		usage = "path to file with urls"
	)

	flag.StringVar(&fileFlag, "file", noFileFlag, usage)
	flag.StringVar(&fileFlag, "f", noFileFlag, usage + " (shorthand)")
}

func checkUrl(url string, ch chan string) {
	answer := fmt.Sprintf("%s - ", url)
	resp, err := http.Get(url)
	if err != nil {
		answer += fmt.Sprintf("Error: %v", err)
	} else {
		defer resp.Body.Close()
		answer += fmt.Sprintf("%v", resp.Status)
	}
	ch <- answer
}


func main() {
	flag.Parse()

	urls, err := urlparser.Parse(fileFlag, flag.Args())
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan string, len(urls))
	for _, url := range urls {
		go checkUrl(url, ch)
	}

	for i := 0; i < len(urls); i++ {
		s := <-ch
		fmt.Println(s)
	}

}