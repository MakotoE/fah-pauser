package main

import (
	"fmt"
	"github.com/MakotoE/go-fahapi"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		const usage = "usage: fah-pauser <path>\nStarts specified program and pauses Folding@home while it is running."
		fmt.Println(usage)
		return
	}

	cmd := exec.Command(os.Args[1])
	if err := cmd.Start(); err != nil {
		log.Panicln(err)
	}

	api, err := fahapi.NewAPI()
	if err != nil {
		log.Panicln(err)
	}

	defer api.Close()

	defer api.UnpauseAll() // Make sure FAH is unpaused in case of panic

	if err := api.PauseAll(); err != nil {
		log.Panicln(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Panicln(err)
	}
}
