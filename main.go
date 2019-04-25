package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execInput(input string) error {
	// remove new lines
	input = strings.TrimSuffix(input, "\n")

	// split string into command and arguments
	args := strings.Split(input, " ")

	// check for built in commands (cd)
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return os.Chdir("/")
		}
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	// prepare command
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		// read text input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
