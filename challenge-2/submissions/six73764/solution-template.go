package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Read input from standard input
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()

		// Call the ReverseString function
		output := ReverseString(input)

		// Print the result
		fmt.Println(output)
	}
}

// ReverseString returns the reversed string of s.
func ReverseString(s string) string {
	run := []rune(s)
	l := len(s)
	for i := 0; i < l/2; i++ {
		run[i], run[l-i-1] = run[l-i-1], run[i]
	}
	return string(run)
}
