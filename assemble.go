package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"strings"
)

// Type for Commands
type command struct {
	number int
	isComment bool
	instruction string
	bytecode string
}

// Error Check
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Starting.")
	fmt.Println()

	// Open File
	inFile, err := os.Open("test.asm") // TODO: Get filename dynamically
	check(err)
	defer inFile.Close()

	// Store in Reader
	reader := bufio.NewReader(inFile)

	// Print Lines
	commands := []command{}
	linecount := 0

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		//// 1st Pass
		// Tokenize Objects
		lineTokens := strings.Split(string(line), " ")
		for i, tok := range(lineTokens) {
			fmt.Printf("%d: %s\t", i, tok)
		}
		fmt.Println()
		// Store Commands
		tcmd := command{number: linecount, isComment: false, instruction: string(line), bytecode: "0x00"}
		commands = append(commands, tcmd)
		linecount++
	}
	fmt.Printf("\n%v\n", commands)

	//// 2nd Pass
	// Through each Command
	// Give an address
	//


	// Close File
	fmt.Println()
	fmt.Println("\nDone.")
}
