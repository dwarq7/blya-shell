package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to BLYA Shell -- Bilyk Ilya Shell!")
	for {
		user, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}

		switch dir {
		case user.HomeDir:
			dir = "~"
		case "/":
			dir = "/"
		default:
			currentDir := strings.Split(dir, "/")
			dir = currentDir[len(currentDir)-1]
		}
		fmt.Print("λ ", dir, " ")

		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Remove the newline character.
		input = strings.TrimSuffix(input, "\n")

		// Skip an empty input.
		if input == "" {
			continue
		}

		// Handle the execution of the input.
		err = execInput(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

// ErrNoPath is returned when 'cd' was called without a second argument.
var ErrNoPath = errors.New("path required")

func execInput(input string) error {
	// Split the input separate the command and the arguments.
	args := strings.Split(input, " ")

	// Check for built-in commands.
	switch args[0] {
	case "":

	case "cd":
		// 'cd' to home with empty path not yet supported.
		if len(args) < 2 {
			return ErrNoPath
		}
		err := os.Chdir(args[1])
		if err != nil {
			return err
		}
		// Stop further processing.
		return nil
	case "exit":
		os.Exit(0)
	case "bye":
		os.Exit(0)
	}

	// Prepare the command to execute.
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and save it's output.
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
