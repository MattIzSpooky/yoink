package main

import (
	"flag"
	"fmt"
	"github.com/MattIzSpooky/yoink/git"
	"io/ioutil"
	"log"
	"net/url"
	"os"
)

func main() {
	if !git.Exists() {
		log.Fatal("An installation of Git is required.")
	}

	fromRepoUrl := flag.String("from", "", "The repository you want to yoink.")
	destinationRepoUrl := flag.String("to", "", "The target repository, where you want the yoink'd repo to end up.")

	flag.Parse()

	if *fromRepoUrl == "" {
		log.Fatal("from flag is required")
	}

	if *destinationRepoUrl == "" {
		log.Fatal("destinationRepoUrl is required")
	}

	_, err := url.ParseRequestURI(*fromRepoUrl)
	if err != nil {
		log.Fatal("Invalid from url")
	}

	_, err = url.ParseRequestURI(*destinationRepoUrl)
	if err != nil {
		log.Fatal("Invalid to url")
	}

	tempDir, err := ioutil.TempDir(os.TempDir(), "yoink.*")

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		fmt.Println(fmt.Sprintf("Cleaning artifacts: %s", tempDir))
		os.RemoveAll(tempDir)
	}()

	gitProgram := git.CreateLocalGit(*fromRepoUrl, *destinationRepoUrl, tempDir)

	if gitProgram.Clone() != nil {
		return
	}

	if gitProgram.ChangeRemote() != nil {
		return
	}

	if gitProgram.Branch() != nil {
		return
	}

	if gitProgram.Push() != nil {
		return
	}
}
