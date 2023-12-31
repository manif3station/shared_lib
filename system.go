package shared_lib

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func System(path string, args ...string) (string, error) {
	cmd := exec.Command(path, args...)
	if os.Getenv("VERBOSE") == "CMD" {
		fmt.Println(">> ", cmd)
	}
	out, err := cmd.Output()
	return string(out), err
}

func SystemRunOnly(path string, args ...string) {
	_, _ = System(path, args...)
}

func SystemOutputOnly(path string, args ...string) string {
	out, _ := System(path, args...)
	return out
}

func Exec(path string, args ...string) error {
	cmd := exec.Command(path, args...)
	if os.Getenv("VERBOSE") == "CMD" {
		fmt.Println(">> ", cmd)
	}
	err := syscall.Exec(cmd.Path, cmd.Args, os.Environ())
	return err
}

func GetArg(n int) string {
	return ArrayItem(os.Args, n)
}
