package main

import (
	"bufio"
	"fmt"
	"os"
)

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan string : ")
	scanner.Scan()
	line := scanner.Text()

	if isPalindrome(line) {
		fmt.Printf("%q merupakan palindrome", line)
	} else {
		fmt.Printf("%q bukan palindrome", line)
	}
}
