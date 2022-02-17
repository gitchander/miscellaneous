package main

import (
	"fmt"
	"io"
	"log"
	"os"

	lrn "github.com/gitchander/miscellaneous/algo/tests/learning"
	"github.com/gitchander/miscellaneous/algo/tests/utils"
)

func main() {
	if true {
		checkError(run(os.Stdin, os.Stdout))
	} else {
		testRW()
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func run(r io.Reader, w io.Writer) error {

	lr := utils.NewLineReader(r)

	blocks, err := utils.LineParseInt(lr)
	if err != nil {
		return err
	}

	for i := 0; i < blocks; i++ {
		result, err := readAndTestBlock(lr)
		if err != nil {
			return err
		}
		fmt.Fprintln(w, result)
	}

	return nil
}

func readAndTestBlock(lr *utils.LineReader) (string, error) {
	n, err := utils.LineParseInt(lr)
	if err != nil {
		return "", err
	}
	var block [2][]lrn.Point
	for i := 0; i < n; i++ {
		xs, err := utils.LineParseIntsN(lr, 3)
		if err != nil {
			return "", err
		}
		p := lrn.Pt(xs[0], xs[1])
		switch t := xs[2]; t {
		case 0:
			block[0] = append(block[0], p)
		case 1:
			block[1] = append(block[1], p)
		default:
			return "", fmt.Errorf("invalid block type %d", t)
		}
	}
	if lrn.SeparableStraight(block[0], block[1]) {
		return "YES", nil
	}
	return "NO", nil
}
