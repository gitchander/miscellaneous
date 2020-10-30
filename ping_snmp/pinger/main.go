package main

import (
	"log"
	"time"

	"github.com/paulstuart/ping"
)

func main() {
	//saveDefaultConfig()
	Main()
}

func Main() {
	var config Config
	err := configLoadToml(&config, "config.toml")
	checkError(err)

	var prevState = -1

	for {
		ok := ping.Ping(config.Ping.Target, config.Ping.Timeout)

		needSet := false
		if ok {
			log.Println("ping success :)")
			if prevState != 1 {
				needSet = true
			}
		} else {
			log.Println("ping fail :(")
			if prevState != 0 {
				needSet = true
			}
		}

		if needSet {
			err = setSnmp(config, ok)
			if err != nil {
				log.Print(err)
			} else {
				if ok {
					prevState = 1
				} else {
					prevState = 0
				}
			}
		}

		time.Sleep(time.Duration(config.Ping.Timeout) * time.Second)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
