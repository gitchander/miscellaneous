package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"math"

	"github.com/gordonklaus/portaudio"
)

func main() {
	err := play()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func test() {

	var buf bytes.Buffer

	//var v int

	sampleRate := 8000
	n := 60 * sampleRate

	for t := 0; t < n; t++ {

		//b := ((t>>6)|t|(t>>uint(t>>16)))*10 + ((t >> 11) & 7)
		//v = (v >> 1) + (v >> 4) + t*(((t>>16)|(t>>6))&(69&(t>>9)))
		//b := (t | (t>>9 | t>>7)) * t & (t>>11 | t>>9)
		//b := (t>>7|t|t>>6)*10 + 4*(t&t>>13|t>>6)

		//b := t * ((t>>12 | t>>8) & 63 & t >> 4)
		//b := (t>>7|t|t>>6)*10 + 4*(t&t>>13|t>>6)

		//
		//b := t*5&(t>>7) | t*3&(t*4>>10)
		//b := (t | (t>>9 | t>>7)) * t & (t>>11 | t>>9)

		b := int(math.Sin(float64(t))) - (t & (t >> 7))

		buf.WriteByte(byte(b))
	}

	err := ioutil.WriteFile("test", buf.Bytes(), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func play() error {

	err := portaudio.Initialize()
	if err != nil {
		return err
	}

	defer portaudio.Terminate()

	var (
		channels         = 1
		samplesPerSecond = 8000
		samplesPerPacket = 100
	)

	var (
		inputChannels   = 0
		outputChannels  = channels
		sampleRate      = float64(samplesPerSecond)
		framesPerBuffer = samplesPerPacket * channels
		out             = make([]uint8, framesPerBuffer)
	)

	stream, err := portaudio.OpenDefaultStream(
		inputChannels,
		outputChannels,
		sampleRate,
		framesPerBuffer,
		&out,
	)
	if err != nil {
		return err
	}

	err = stream.Start()
	if err != nil {
		return err
	}
	defer stream.Stop()

	var t int
	for i := 0; i < 10000; i++ {

		for j := 0; j < framesPerBuffer; j++ {

			b := (t | (t>>9 | t>>7)) * t & (t>>11 | t>>9)

			t++
			out[j] = uint8(b)
		}

		err = stream.Write()
		if err != nil {
			return err
		}
	}

	return nil
}
