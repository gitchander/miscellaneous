package main

import (
	"fmt"
	"sort"
	"time"

	sorta "github.com/gitchander/miscellaneous/sort_algo"
	"github.com/gitchander/miscellaneous/sort_algo/util"
)

func main() {
	baseExample()
	//compareQuickAndMerge()
}

func baseExample() {

	a := make([]int, 1000000)
	for i := range a {
		a[i] = i + 1
	}

	data := sort.IntSlice(a)

	//r := newRandSeed(0)
	r := util.RandomerTime(time.Now())
	util.Scramble(r, data)
	//reverse(data)

	//fmt.Println(a)

	start := time.Now()
	//sort.Sort(data)
	//sorta.BubbleSort(data)
	//sorta.CocktailSort(data)
	//sorta.InsertionSort(data)
	//sorta.SelectionSort(data)
	//sorta.GnomeSort(data)
	//sorta.StoogeSort(data)
	//sorta.ShellSort(data)
	sorta.CombSort(data)
	//sorta.QuickSort(data)
	dur := time.Since(start)

	fmt.Println("QuickSort:")
	fmt.Println(dur)
	printSorted(data)
	fmt.Println()
}

func compareQuickAndMerge() {
	quickSortExample(1000000, true)
	mergeSortExample(1000000, true)

	fmt.Println("--------------------")

	quickSortExample(10000, false)
	mergeSortExample(10000, false)
}

func quickSortExample(n int, scramble bool) {
	a := make([]int, n)
	for i := range a {
		a[i] = i + 1
	}
	data := sort.IntSlice(a)
	if scramble {
		util.Scramble(util.RandomerTimeNow(), data)
	}

	start := time.Now()
	sorta.QuickSort(data) // <- sort!
	dur := time.Since(start)

	fmt.Println("QuickSort:")
	fmt.Println(dur)
	printSorted(data)
	fmt.Println()
}

func mergeSortExample(n int, scramble bool) {
	a := make([]int, n)
	for i := range a {
		a[i] = i + 1
	}
	data := sort.IntSlice(a)
	if scramble {
		util.Scramble(util.RandomerTimeNow(), data)
	}

	start := time.Now()
	mergeSortInts(a) // <- sort!
	dur := time.Since(start)

	fmt.Println("MergeSort:")
	fmt.Println(dur)
	printSorted(data)
	fmt.Println()
}

func mergeSortInts(a []int) {
	var (
		n = len(a)
		b = make([]int, 2*n)
	)
	copy(b, a)
	sorta.MergeSort(sort.IntSlice(b), n)
	copy(a, b[:n])
}

func printSorted(data sort.Interface) {
	if sort.IsSorted(data) {
		fmt.Println("Success! Data is sorted.")
	} else {
		fmt.Println("Unsuccess! Data isn't sorted.")
	}
}
