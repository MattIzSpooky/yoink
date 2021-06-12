package git

import (
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
}

func Exists() bool {
	_, err := exec.LookPath("git")

	return err == nil
}

func (l LocalGit) Clone() error {
	_, err := exec.Command("git", "clone", l.From, l.Folder).Output()
	return err
}

func (l LocalGit) ChangeRemote() error {
	_, err := exec.Command("git", "-C", l.Folder, "remote", "set-url", "origin", l.To).Output()
	return err
}

func (l LocalGit) Branch() error {
	_, err := exec.Command("git", "-C", l.Folder, "branch", "-M", "main").Output()
	return err
}

func (l LocalGit) Push() error {
	_, err := exec.Command("git", "-C", l.Folder, "push", "-u", "origin", "main").Output()
	return err
}