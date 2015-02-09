// +build !windows

package main

import (
	"os/exec"
)

func playsnd(filename string) error {
	return exec.Command("open", filename).Run()
}
