package main

import (
	"flag"
	"io"
	"log"
	"os"
	"time"

	"github.com/gordonklaus/portaudio"
)

func main() {

	var config Config

	flag.IntVar(&(config.Channels), "channels", 1, "channels per sample")
	flag.IntVar(&(config.SamplesPerSecond), "rate", 8000, "sample rate")
	flag.StringVar(&(config.Format), "format", "", "sample format")

	flag.Parse()

	err := Run(config)
	checkError(err)
}

type Config struct {
	Channels         int
	SamplesPerSecond int
	Format           string // u8, i8, u16_le, u16be, ...
}

var DefaultConfig = Config{
	Channels:         1,
	SamplesPerSecond: 8000,
}

func getBufferDuration(samplesPerSecond, samplesPerBuffer int) time.Duration {
	return (time.Second * time.Duration(samplesPerBuffer)) / time.Duration(samplesPerSecond)
}

func getSamplesPerBuffer(samplesPerSecond int, bufferDuration time.Duration) int {
	return int((bufferDuration * time.Duration(samplesPerSecond)) / time.Second)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Run(config Config) error {

	err := portaudio.Initialize()
	if err != nil {
		return err
	}
	defer portaudio.Terminate()

	var (
		bufferDuration   = 20 * time.Millisecond
		samplesPerBuffer = getSamplesPerBuffer(config.SamplesPerSecond, bufferDuration)
		framesPerBuffer  = samplesPerBuffer * config.Channels
	)

	var (
		inputChannels  = 0
		outputChannels = config.Channels
		sampleRate     = float64(config.SamplesPerSecond)
		buffer         = make([]byte, framesPerBuffer)
	)

	stream, err := portaudio.OpenDefaultStream(
		inputChannels,
		outputChannels,
		sampleRate,
		framesPerBuffer,
		&buffer,
	)
	if err != nil {
		return err
	}

	err = stream.Start()
	if err != nil {
		return err
	}
	defer stream.Stop()

	for {
		_, err := io.ReadFull(os.Stdin, buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		err = stream.Write()
		if err != nil {
			return err
		}
	}
	return nil
}
