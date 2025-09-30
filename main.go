package main

import (
	"fmt"
)

func main() {
	json := `"url": "https://example.com", "enabled": true`
	fmt.Println(json)
}

type Request struct {
	URL    string
	Enable bool
}

func Tokenize(input string) []string {
	var tokenStr []string

	return tokenStr
}
