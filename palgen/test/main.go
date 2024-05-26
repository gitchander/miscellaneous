package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"

	"github.com/gitchander/miscellaneous/palgen"
	"github.com/gitchander/miscellaneous/utils/random"
)

func main() {
	size := image.Point{512, 64}
	m := NewRGBASize(size)

	var p palgen.Params
	palgen.RandParams(random.NewRandNow(), &p, 7)
	printJSON(p)

	palgen.DrawPalette(m, p)

	err := SaveImagePNG("palette.png", m)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func printJSON(v interface{}) {
	data, err := json.MarshalIndent(v, "", "\t")
	checkError(err)
	fmt.Printf("%s\n", data)
}

func NewRGBASize(size image.Point) *image.RGBA {
	r := image.Rectangle{Max: size}
	return image.NewRGBA(r)
}

func SaveImagePNG(filename string, m image.Image) error {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, buf.Bytes(), 0666)
}
