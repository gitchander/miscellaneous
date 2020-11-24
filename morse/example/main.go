package main

import (
	"fmt"
	"log"

	"github.com/gitchander/miscellaneous/morse"
)

func main() {
	samples := []string{
		"SOS",
		"MORSE CODE",
	}
	for _, sample := range samples {
		data, err := morse.Units(sample)
		checkError(err)
		fmt.Println(string(data))
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
