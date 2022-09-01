package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/gitchander/miscellaneous/crypto/sha256test/core"
)

func main() {
	checkError(run(os.Args[1:]))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func run(fileNames []string) error {

	if len(fileNames) == 0 {
		var b bytes.Buffer
		_, err := io.Copy(&b, os.Stdin)
		if err != nil {
			return err
		}
		return runBytes(b.Bytes())
	}

	for _, fileName := range fileNames {
		bs, err := ioutil.ReadFile(fileName)
		if err != nil {
			return err
		}
		err = runBytes(bs)
		if err != nil {
			return err
		}
	}

	return nil
}

func runBytes(bs []byte) error {

	var r core.Result
	err := json.Unmarshal(bs, &r)
	if err != nil {
		return err
	}

	s, err := core.CalcResult(r)
	if err != nil {
		return err
	}

	fmt.Println(s)

	return nil
}
