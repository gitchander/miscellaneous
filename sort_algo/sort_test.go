package sort_algo

import (
	"sort"
	"testing"

	. "github.com/gitchander/miscellaneous/sort_algo/util"
)

func TestMergeSort(t *testing.T) {

	//r := RandomerSeed(0)
	r := RandomerTimeNow()

	a := make([]int, 0, 10000)

	for i := 0; i < 1000; i++ {

		n := r.Intn(cap(a) / 2)

		a = a[:n]
		for i := range a {
			a[i] = i + 1
		}

		// temp
		b := a[n : 2*n]
		for i := range b {
			b[i] = 0
		}

		data := sort.IntSlice(a[:2*n])

		ScrambleN(r, data, n)

		MergeSort(data, n)

		if !sort.IsSorted(sort.IntSlice(a)) {
			t.Fatal("data isn't sorted")
		}

		for i := range a {
			if x := i + 1; a[i] != x {
				t.Fatalf("data: ([%d]:%d) != %d", i, a[i], x)
			}
		}
		for i := range b {
			if b[i] != 0 {
				t.Fatalf("temp: ([%d]:%d) != %d", i, b[i], 0)
			}
		}
	}
}

func TestMergeSortLarge(t *testing.T) {

	const n = 1000000
	a := make([]int, n, 2*n)
	for i := range a {
		a[i] = i + 1
	}

	data := sort.IntSlice(a[:2*n])

	//r := RandomerSeed(0)
	r := RandomerTimeNow()
	ScrambleN(r, data, n)

	MergeSort(data, n)

	if !sort.IsSorted(sort.IntSlice(a)) {
		t.Fatal("data isn't sorted")
	}

	for i := range a {
		if x := i + 1; a[i] != x {
			t.Fatalf("data: ([%d]:%d) != %d", i, a[i], x)
		}
	}
}
