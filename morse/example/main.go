package main

import (
	"fmt"
	"log"

	"github.com/gitchander/miscellaneous/morse"
)

func main() {

	text := "SOS"

	data, err := morse.Units(text)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
