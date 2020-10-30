package main

import (
	"errors"
	"fmt"

	"github.com/gitchander/miscellaneous/ujson"
)

type Mammal struct{}

func (*Mammal) isAnimal_Value() {}

type Bird struct {
	Name string
}

func (*Bird) isAnimal_Value() {}

type Fish struct{}

func (*Fish) isAnimal_Value() {}

type Reptile struct{}

func (*Reptile) isAnimal_Value() {}

type Amphibian struct{}

func (*Amphibian) isAnimal_Value() {}

type Arthropod struct{}

func (*Arthropod) isAnimal_Value() {}

const (
	key_Mammal    = "mammal"
	key_Bird      = "bird"
	key_Fish      = "fish"
	key_Reptile   = "reptile"
	key_Amphibian = "amphibian"
	key_Arthropod = "arthropod"
)

func getAnimalKey(val isAnimal_Value) (key string, err error) {
	switch val.(type) {
	case *Mammal:
		key = key_Mammal
	case *Bird:
		key = key_Bird
	case *Fish:
		key = key_Fish
	case *Reptile:
		key = key_Reptile
	case *Amphibian:
		key = key_Amphibian
	case *Arthropod:
		key = key_Arthropod
	default:
		err = errors.New("invalid animal value")
	}
	return
}

func makeAnimal(key string) (val isAnimal_Value, err error) {
	switch key {
	case key_Mammal:
		val = new(Mammal)
	case key_Bird:
		val = new(Bird)
	case key_Fish:
		val = new(Fish)
	case key_Reptile:
		val = new(Reptile)
	case key_Amphibian:
		val = new(Amphibian)
	case key_Arthropod:
		val = new(Arthropod)
	default:
		err = fmt.Errorf("invalid animal [%s]", key)
	}
	return
}

func makeValue_Animal(key string) (interface{}, error) {
	return makeAnimal(key)
}

type isAnimal_Value interface {
	isAnimal_Value()
}

type Animal struct {
	Value isAnimal_Value
}

func (a Animal) MarshalJSON() ([]byte, error) {
	key, err := getAnimalKey(a.Value)
	if err != nil {
		return nil, err
	}
	return ujson.Marshal(key, a.Value)
}

func (a *Animal) UnmarshalJSON(data []byte) error {
	val, err := ujson.Unmarshal(data, makeValue_Animal)
	if err != nil {
		return err
	}
	a.Value = val.(isAnimal_Value)
	return nil
}
