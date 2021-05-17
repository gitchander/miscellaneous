package main

import (
	"flag"
	"fmt"
	"os"
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

	lr := NewLineReader(os.Stdin)

	var stat Stat

	var tasker Tasker = newMathTasker(r, bos)
	var task Task

	for {
		if task == nil {
			newTask, ok := tasker.Next()
			if !ok {
				break
			}
			task = newTask
		}

		fmt.Print(task.Question())

		line, err := lr.ReadLine()
		if err != nil {
			return err
		}

		ok, failMessage := task.CheckAnswer(line)
		if !ok {

			// 	wrongPhrase := randWrongPhrase(r)
			// 	fmt.Printf("%s: have %d, want %d\n", wrongPhrase, result, z)

			fmt.Println(failMessage)
			stat.Add(false)
			fmt.Println("Statistics:", stat)
			fmt.Println()

			continue
		}

		fmt.Println(randCorrectPhrase(r))
		stat.Add(true)
		fmt.Println("Statistics:", stat)
		fmt.Println()

		task = nil // reset task
	}

	return nil
}
