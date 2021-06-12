package git

import (
	"fmt"
	"github.com/gookit/color"
	"log"
	"os"
	"os/exec"
)

type Program interface {
	Clone() error
	ChangeRemote() error
	Branch() error
	Push() error
}

type LocalGit struct {
	From   string
	To     string
	Folder string

	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func CreateLocalGit(from string, to string, folder string) Program {
	return LocalGit{
		From:       from,
		To:          to,
		Folder:      folder,
		infoLogger:  log.New(os.Stdout, color.Green.Sprint("[Info]: "), log.Ldate|log.Ltime),
		errorLogger: log.New(os.Stderr, color.Red.Sprint("[Error]: "), log.Ldate|log.Ltime),
	}
}

func Exists() bool {
	_, err := exec.LookPath("git")

	return err == nil
}

func (l LocalGit) Clone() error {
	_, err := exec.Command("git", "clone", l.From, l.Folder).Output()

	if err != nil {
		l.handleError(err)
		return err
	}

	l.infoLogger.Println(fmt.Sprintf("Succesfully cloned %s in %s", l.From, l.Folder))
	return nil
}

func (l LocalGit) ChangeRemote() error {
	_, err := exec.Command("git", "-C", l.Folder, "remote", "set-url", "origin", l.To).Output()

	if err != nil {
		l.handleError(err)
		return err
	}
	l.infoLogger.Println(fmt.Sprintf("Succesfully changed remote from %s", l.From))
	l.infoLogger.Println(fmt.Sprintf("Succesfully changed remote to %s", l.To))

	return nil
}

func (l LocalGit) Branch() error {
	_, err := exec.Command("git", "-C", l.Folder, "branch", "-M", "main").Output()

	if err != nil {
		l.handleError(err)
		return err
	}

	l.infoLogger.Println("Successfully branched")

	return nil
}

func (l LocalGit) Push() error {
	l.infoLogger.Println("Pushing code... Might take a while")

	_, err := exec.Command("git", "-C", l.Folder, "push", "-u", "origin", "main").Output()

	if err != nil {
		l.handleError(err)
		return err
	}

	l.infoLogger.Println("Done!")
	l.infoLogger.Println(fmt.Sprintf("Go to: %s", l.To))

	return nil
}

func (l LocalGit) handleError(err error) {
	l.errorLogger.Println(fmt.Sprintf("Git returned: %s", err.Error()))
}