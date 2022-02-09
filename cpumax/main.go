package main

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {

	n := runtime.NumCPU()

	for i := 0; i < n; i++ {
		go run()
	}

	waitSignalOS()
}

func run() {
	for {
	}
}

func waitSignalOS() os.Signal {
	ls := make(chan os.Signal, 1)
	signal.Notify(ls, syscall.SIGINT, syscall.SIGTERM)
	l := <-ls
	return l
}
