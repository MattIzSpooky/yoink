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
	defer os.RemoveAll(tempDir)

	gitProgram := git.LocalGit{
		From:   *fromRepoUrl,
		To:     *destinationRepoUrl,
		Folder: tempDir,
	}

	err = gitProgram.Clone()
	if err != nil {
		log.Println(fmt.Sprintf("Git returned: %s", err.Error()))
		return
	}
	fmt.Println(fmt.Sprintf("Succesfully cloned %s to system temp", *fromRepoUrl))

	err = gitProgram.ChangeRemote()
	if err != nil {
		log.Println(fmt.Sprintf("Git returned: %s", err.Error()))
		return
	}
	fmt.Println(fmt.Sprintf("Succesfully changed remote from %s", *fromRepoUrl))
	fmt.Println(fmt.Sprintf("Succesfully changed remote to %s", *destinationRepoUrl))

	err = gitProgram.Branch()
	if err != nil {
		log.Println(fmt.Sprintf("Git returned: %s", err.Error()))
		return
	}

	err = gitProgram.Push()
	if err != nil {
		log.Println(fmt.Sprintf("Git returned: %s", err.Error()))
		return
	}
	fmt.Println("Done!")
}
