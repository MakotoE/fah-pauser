package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		const usage = "usage: fah-pauser <path>\nStops Folding@home when <path> is running"
		fmt.Println(usage)
		return
	}

	cmd := exec.Command(os.Args[1])
	if err := cmd.Start(); err != nil {
		log.Panicln(err)
	}

	api, err := NewAPI()
	if err != nil {
		log.Panicln(err)
	}

	defer api.Close()

	defer api.Unpause() // Make sure FAH is unpaused in case of panic

	if err := api.Pause(); err != nil {
		log.Panicln(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Panicln(err)
	}
}