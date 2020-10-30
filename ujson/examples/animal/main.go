package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	exampleAnimal()
	exampleParseAnimal()
}

func exampleAnimal() {
	a := Animal{
		Value: &Amphibian{},
		//Value: &Mammal{},
		//		Value: &Bird{
		//			Name: "Bird",
		//		},
	}
	data, err := json.Marshal(a)
	checkError(err)

	fmt.Println(string(data))

	var b Animal

	err = json.Unmarshal(data, &b)
	checkError(err)

	//data, err = json.Marshal(b)
	data, err = json.MarshalIndent(b, "", "\t")
	checkError(err)

	fmt.Println(string(data))
}

func exampleParseAnimal() {

	text := `{"bird":{"Name":"hello"}}`
	data := []byte(text)

	var b Animal

	err := json.Unmarshal(data, &b)
	checkError(err)

	//data, err = json.Marshal(b)
	data, err = json.MarshalIndent(b, "", "\t")
	checkError(err)

	fmt.Println(string(data))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
