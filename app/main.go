package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/vadyaov/url_access/urlparser"
)

// 1. ./app https://google.com https://example.com http://nonexistent-site.org
// 2. ./my_link_checker -file urls.txt

var fileFlag string

func init() {
	const (
		noFileFlag = ""
		usage = "path to file with urls"
	)

	flag.StringVar(&fileFlag, "file", noFileFlag, usage)
	flag.StringVar(&fileFlag, "f", noFileFlag, usage + " (shorthand)")
}


func main() {
	flag.Parse()


	urls, err := urlparser.Parse(fileFlag, flag.Args())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(urls)

}