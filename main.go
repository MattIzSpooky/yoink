package main

import (
	"flag"
	"log"
	"os/exec"
)

func main() {
	gitPath, err := exec.LookPath("git")
	if err != nil {
		log.Fatal("An installation of Git is required.")
	}

	log.Println(gitPath)

	fromRepo := flag.String("from", "", "The repository you want to yoink.")
	destination := flag.String("to", "", "The target repository, where you want the yoink'd repo to end up.")

	flag.Parse()

	if *fromRepo == "" {
		log.Fatal("from flag is required")
		return
	}

	if *destination == "" {
		log.Fatal("destination is required")
		return
	}

	// TODO: Validate inputs
	// TODO: Only support Github because easiest

	log.Println(*fromRepo)
	log.Println(*destination)
}
