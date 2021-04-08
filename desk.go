package main

import (
	"bytes"
	"os/exec"
)

// This will be the desk interface.  Exclusively interaface with desks/idasen for now (TODO)

type Desk interface {
	To(position string) error
	Status() (string, error)
	Save(position string) error
}

// This is the idasen module
type Idasen struct {
}

func (d *Idasen) To(position string) error {
	cmd := exec.Command("idasen", position)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return err
}

func (d *Idasen) Status() (string, error) {
	cmd := exec.Command("idasen", "height")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

func (d *Idasen) Save(position string) error {
	cmd := exec.Command("idasen", "save", position)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return err
}
