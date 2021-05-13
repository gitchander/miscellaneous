package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var c Config

	flag.StringVar(&(c.Operations), "ops", "add,sub,mul,div", "operations")

	flag.Parse()

	err := run(c)
	checkError(err)
}

type Config struct {
	Operations string
}

func run(c Config) error {

	bos, err := parseOperations(c.Operations)
	if err != nil {
		return err
	}

	r := newRandNow()

	br := bufio.NewReader(os.Stdin)

	var stat Stat

	var bo BinaryOperation
	var x, y int // Arguments
	lastCorrect := true

	for {
		if lastCorrect {
			bo = randBinaryOperation(r, bos)
			x, y = randArguments(r, bo)
		}

		fmt.Printf("%d %s %d = ", x, bo.Symbol(), y)

		line, isPrefix, err := br.ReadLine()
		if err != nil {
			return err
		}

		if false {
			fmt.Printf("user input: line: %q, isPrefix: %t\n", string(line), isPrefix)
		}

		result, err := strconv.Atoi(string(line))
		if err != nil {
			//stat.Add(false)
			fmt.Printf("invalid input data: %q\n", string(line))
			fmt.Println()
			lastCorrect = false
			continue
		}

		z := bo.Do(x, y)
		if result != z {
			stat.Add(false)
			wrongPhrase := randWrongPhrase(r)
			fmt.Printf("%s: have %d, want %d\n", wrongPhrase, result, z)
			lastCorrect = false
		} else {
			stat.Add(true)
			fmt.Println(randCorrectPhrase(r))
			lastCorrect = true
		}

		fmt.Println("Statistics:", stat)
		fmt.Println()
	}
	return nil
}
