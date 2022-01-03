package main

import (
	"console"
	"log"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		command := os.Args[1]
		defaults := map[string]string{"workspace": "default"}
		flags := console.ParseFlags(os.Args[2:], defaults)

		switch command {
		case "clean":
			scope, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
				break
			}
			err = cleanScope(scope, flags["workspace"])
			if err != nil {
				log.Fatal(err)
				break
			}
		case "help", "-help", "--help":
			usage()
		default:
			console.Yellow("unrecognized command: " + command + "\n")
		}
	} else {
		usage()
	}
}

func usage() {
	console.White("Usage: tf COMMAND [args]\n\n")
	console.White("commands:\n")
	console.White("  clean - wipes terraform cache from current scope, and re-inits terraform\n")
	console.White("args:\n")
	console.White("  -workspace=<workspace to select after initialization>\n")
}
