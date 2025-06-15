package urlparser

import (
	"bufio"
	"fmt"
	"os"
)

func Parse(file string, args []string) ([]string, error) {
	if file != "" {
		fmt.Printf("Parsing URLs from file: %v\n", file)
		return fromFile(file)
	}

	fmt.Println("Parsing URLs from command-line arguments")
	return args, nil
}

func fromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			urls = append(urls, line)
		}
	}
	return urls, scanner.Err()
}