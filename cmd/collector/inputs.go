package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadContentFromInput() string {
	fmt.Println("Enter content (finish with an empty line):")

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}
