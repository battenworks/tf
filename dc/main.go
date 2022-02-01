package main

import (
	"console"
	"log"
	"os"
)

type Executor interface {
	Execute(cmdName string, cmdArgs []string) ([]byte, error)
}

var executor Executor

func main() {
	if len(os.Args) > 1 {
		command := os.Args[1]

		switch command {
		case "up":
			err := upDisconnected(executor)
			if err != nil {
				log.Fatal(err)
				break
			}
			if err != nil {
				log.Fatal(err)
				break
			}
		case "help", "-help", "--help":
			usage()
		default:
			err := passThrough(executor, os.Args[1:])
			if err != nil {
				log.Fatal(err)
				break
			}
		}
	} else {
		usage()
	}
}

func usage() {
	console.White("Usage: dc COMMAND\n\n")
}
