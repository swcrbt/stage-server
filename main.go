package main

import (
	"log"

	"github.com/swcrbt/stage-server/cmd"
)

func main() {
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}