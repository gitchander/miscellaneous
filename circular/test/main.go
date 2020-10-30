package main

import (
	"fmt"

	"github.com/gitchander/miscellaneous/circular"
)

func main() {
	//var cb circular.Buffer
	var cb = circular.NewBuffer(10)

	fmt.Println("empty:", cb.Empty())
	fmt.Println("full:", cb.Full())

	data := []byte{0, 1, 2, 3, 4, 5, 6}

	n, err := cb.Write(data[:])
	if err != nil {
		fmt.Printf("%d, %s\n", n, err)
	}

	tmp := make([]byte, 100)
	for i := 0; i < 3; i++ {
		n, err = cb.Read(tmp)
		if err != nil {
			fmt.Printf("%d, %s\n", n, err)
		}
		fmt.Printf("read: %d, data: [% X]\n", n, tmp[:n])
	}
}
