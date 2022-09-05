package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gitchander/miscellaneous/crypto/sha256test/core"
)

func main() {
	var c Config
	flag.StringVar(&(c.Text), "text", "Hello, World!", "source text")
	flag.IntVar(&(c.SampleSize), "sample_size", 4, "sample size")
	flag.IntVar(&(c.OriginSize), "origin_size", 8, "origin size")
	flag.Parse()
	checkError(run(c))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Config struct {
	Text       string
	SampleSize int
	OriginSize int
}

func run(c Config) error {

	start := time.Now()

	ctx := handleSignalOS(context.Background())

	result, err := core.Search(ctx, []byte(c.Text), c.SampleSize, c.OriginSize)
	if err != nil {
		return err
	}

	printJSON(result)

	fmt.Println("work duration:", time.Since(start))

	return nil
}

func printJSON(v interface{}) {
	data, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func handleSignalOS(ctx context.Context) context.Context {

	child, cancel := context.WithCancel(ctx)

	signals := make(chan os.Signal, 1)

	signal.Notify(signals,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	go func() {
		for {
			sig := <-signals
			switch sig {
			case syscall.SIGHUP:
				fmt.Println("there is signal SIGHUP")
				cancel()
				return
			case syscall.SIGINT:
				fmt.Println("there is signal SIGINT")
				cancel()
				return
			case syscall.SIGTERM:
				fmt.Println("there is signal SIGTERM")
				cancel()
				return
			}
		}
	}()

	return child
}
