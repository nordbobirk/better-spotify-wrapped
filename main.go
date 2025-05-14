package main

import (
	"better-spotify-wrapped/cli"
	"better-spotify-wrapped/internal"
	"log"
	"os"
)

func main() {
	initDataDir()
	cli.Run(os.Args)
}

func initDataDir() {
	if !dataDirExists() {
		err := os.Mkdir(internal.DataDir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func dataDirExists() bool {
	_, err := os.Stat(internal.DataDir)
	return err == nil || !os.IsNotExist(err)
}
