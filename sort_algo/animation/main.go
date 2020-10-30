package main

import (
	"bytes"
	"image/gif"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"

	sorta "github.com/gitchander/miscellaneous/sort_algo"
	"github.com/gitchander/miscellaneous/sort_algo/util"
)

const dir = "result"

func main() {
	err := os.Mkdir(dir, 0777)
	if !os.IsExist(err) {
		checkError(err)
	}

	//renderOne()
	//renderMergeSort()
	renderAll()
}

func renderOne() {
	params{
		n:        50,
		delay:    10,
		width:    10,
		f:        sorta.ShellSort,
		filename: filepath.Join(dir, "test.gif"),
	}.run()
}

func renderAll() {

	// BubbleSort
	params{
		n:        20,
		delay:    10,
		width:    10,
		f:        sorta.BubbleSort,
		filename: filepath.Join(dir, "bubble-sort.gif"),
	}.run()

	// CocktailSort
	params{
		n:        20,
		delay:    10,
		width:    10,
		f:        sorta.CocktailSort,
		filename: filepath.Join(dir, "cocktail-sort.gif"),
	}.run()

	// InsertionSort
	params{
		n:        50,
		delay:    10,
		width:    10,
		f:        sorta.InsertionSort,
		filename: filepath.Join(dir, "insertion-sort.gif"),
	}.run()

	// SelectionSort
	params{
		n:        20,
		delay:    10,
		width:    10,
		f:        sorta.SelectionSort,
		filename: filepath.Join(dir, "selection-sort.gif"),
	}.run()

	// GnomeSort
	params{
		n:        20,
		delay:    10,
		width:    10,
		f:        sorta.GnomeSort,
		filename: filepath.Join(dir, "gnome-sort.gif"),
	}.run()

	// StoogeSort
	params{
		n:        10,
		delay:    10,
		width:    10,
		f:        sorta.StoogeSort,
		filename: filepath.Join(dir, "stooge-sort.gif"),
	}.run()

	// CombSort
	params{
		n:        50,
		delay:    10,
		width:    10,
		f:        sorta.CombSort,
		filename: filepath.Join(dir, "comb-sort.gif"),
	}.run()

	// ShellSort
	params{
		n:        50,
		delay:    10,
		width:    10,
		f:        sorta.ShellSort,
		filename: filepath.Join(dir, "shell-sort.gif"),
	}.run()

	// QuickSort
	params{
		n:        50,
		delay:    10,
		width:    10,
		f:        sorta.QuickSort,
		filename: filepath.Join(dir, "quick-sort.gif"),
	}.run()

	// MergeSort
	renderMergeSort()
}

func renderMergeSort() {
	p := params{
		n:        50,
		delay:    10,
		width:    4,
		f:        nil,
		filename: filepath.Join(dir, "merge-sort.gif"),
	}
	mergeSort(p)
}

func makeGIF(filename string, render func(*gif.GIF) error) error {
	anim := new(gif.GIF)
	err := render(anim)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	err = gif.EncodeAll(&buf, anim)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, buf.Bytes(), 0666)
	return err
}

func baseSort(p params) {
	vs := make([]int, p.n)
	for i := range vs {
		vs[i] = i
	}
	util.Scramble(util.RandomerTimeNow(), sort.IntSlice(vs))

	render := func(anim *gif.GIF) error {
		d := NewDrawer(p, vs, anim)
		data := NewIntSlice(vs, d)
		p.f(data)
		if !dataIsSorted(data, data.Len()) {
			log.Fatal("data isn't sorted")
		}
		return nil
	}

	err := makeGIF(p.filename, render)
	checkError(err)
}

func mergeSort(p params) {
	vs := make([]int, p.n)
	for i := range vs {
		vs[i] = i
	}
	util.Scramble(util.RandomerTimeNow(), sort.IntSlice(vs))

	render := func(anim *gif.GIF) error {
		n := len(vs)
		ws := make([]int, 2*n)
		copy(ws, vs)
		d := NewDrawer(p, ws, anim)
		data := NewIntSlice(ws, d)
		sorta.MergeSort(data, n)
		copy(vs, ws)
		if !dataIsSorted(data, n) {
			log.Fatal("data isn't sorted")
		}
		return nil
	}

	err := makeGIF(p.filename, render)
	checkError(err)
}

func dataIsSorted(data sort.Interface, n int) bool {
	for n > 1 {
		n--
		if data.Less(n, n-1) {
			return false
		}
	}
	return true
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
