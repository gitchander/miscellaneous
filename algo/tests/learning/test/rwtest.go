package main

import (
	"fmt"
	"strings"
)

func testRW() {

	inputText := `2
6
1 1 1
1 2 1
1 3 0
2 1 1
2 2 0
3 1 0
6
1 3 0
2 2 0
1 2 1
3 1 1
2 1 1
1 1 0`

	outputText := `YES
NO
`

	r := strings.NewReader(inputText)

	var b strings.Builder
	checkError(run(r, &b))
	fmt.Print(b.String())

	if b.String() != outputText {
		checkError(fmt.Errorf("invalid result"))
	}
}
