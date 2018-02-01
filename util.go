package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func askYesNo(msg string, defaultYes bool) bool {
	reader := bufio.NewReader(os.Stdin)
	prompt := "[Y/n]"
	if !defaultYes {
		prompt = "[y/N]"
	}
	fmt.Printf("%s %s ", msg, prompt)
	line, err := reader.ReadByte()
	if err != nil || string(line) == "\n" {
		return defaultYes
	}
	low := strings.ToLower(string(line))
	if low == "y" {
		return true
	} else if low == "n" {
		return false
	} else {
		fmt.Println("I didn't quite get that.")
		return askYesNo(msg, defaultYes)
	}
}

// Util to run a command and pipe output to the current terminal
func runCmd(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Util to install packages with apt
func aptInstall(packages ...string) error {
	args := []string{"install", "-y"}
	args = append(args, packages...)
	return runCmd("apt-get", args...)
}

// Copies src to dst, overwriting dst if it exists
func copyOver(src, dst string, askToOverwrite bool) error {
	_, err := os.Stat(dst)
	if err == nil && askToOverwrite {
		msg := fmt.Sprintf("File already exists at %s, overwrite it?", dst)
		if askYesNo(msg, true) {
			os.Remove(dst)
		} else {
			return nil
		}
	}
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	err = os.MkdirAll(filepath.Dir(dst), 0755)
	if err != nil {
		return err
	}
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Sync()
}
