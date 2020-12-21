package main

import (
	"fmt"
	"time"

	"github.com/gitchander/miscellaneous/algo/math/partition"
)

func main() {
	//testOne()
	testOneWalk()
	testRange()
	//testRangeTest()
}

func testOne() {
	start := time.Now()
	n := 100
	p := partition.Partition(n)
	fmt.Printf("p(%d) = %d\n", n, p)
	fmt.Println(time.Since(start))
}

func testOneWalk() {

	n := 5

	f := func(as []int) {
		fmt.Println(as)
	}

	p := partition.PartitionWalk(n, f)
	fmt.Printf("p(%d) = %d\n", n, p)
}

func testRange() {
	for n := -5; n < 31; n++ {
		p := partition.Partition(n)
		fmt.Printf("p(%d) = %d\n", n, p)
	}
}

func testRangeTest() {
	if true {
		f := func(as []int) {
			fmt.Println(as)
		}
		n := 4
		p := partition.PartitionTest(n, f)
		fmt.Printf("p(%d) = %d\n", n, p)
	}
	if true {
		f := func([]int) {}
		for n := -5; n < 31; n++ {
			p := partition.PartitionTest(n, f)
			fmt.Printf("p(%d) = %d\n", n, p)
		}
	}
}
