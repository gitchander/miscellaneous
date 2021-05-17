package main

import (
	"bufio"
	"io"
	"log"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func lerp(v0, v1 float64, t float64) float64 {
	return v0*(1-t) + v1*t
}

func percent(part, full float64) float64 {
	return part * 100 / full
}

//------------------------------------------------------------------------------
type LineReader struct {
	br *bufio.Reader
}

func NewLineReader(r io.Reader) *LineReader {
	return &LineReader{
		br: bufio.NewReader(r),
	}
}

func (p *LineReader) ReadLine() (string, error) {
	line, _, err := p.br.ReadLine()
	if err != nil {
		return "", err
	}
	return string(line), nil
}
