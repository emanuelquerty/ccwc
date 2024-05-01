package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/emanuelquerty/ccwc"
)

func main() {
	args := os.Args[1:]
	ccwc.DefineFlag("m", "outputs the number of chars")
	ccwc.DefineFlag("c", "outputs the number of bytes")
	ccwc.DefineFlag("w", "outputs the number of words")
	ccwc.DefineFlag("l", "outputs the number of lines")

	ccwc.SetFlag(args)
	flagValue, ok := ccwc.GetFlag()
	
	var reader io.Reader
	switch {
	case ok && flagValue.Value != "": // user specifies a flag with filepath
		reader = openFile(flagValue.Value)
	case !ok && len(args) > 0: // user does not specifies any flag but filepath
		reader = openFile(args[0])
	default: // user does not specify filepath
		reader = bufio.NewReader(os.Stdin)
	}

	bytes, err := ccwc.ReadFile(reader)
	if err != nil {
		log.Fatal(err)
	}

	var count int
	switch flagValue.Name {
	case "c":
		count = ccwc.FindByteCount(bytes)
	case "m":
		count = ccwc.FindCharacterCount(bytes)
	case "w":
		count = ccwc.FindWordCount(bytes)
	case "l":
		count = ccwc.FindLineCount(bytes)
	case "":
		byteCount, lineCount, wordCount := findAll(bytes)
		fmt.Println(lineCount, wordCount, byteCount, flagValue.Value)
		return
	}
	fmt.Println("  ", count, flagValue.Value)
}

func findAll(fileBytes []byte) (byteCount int, lineCount int, wordCount int) {
	byteCount = ccwc.FindByteCount(fileBytes)
	lineCount = ccwc.FindLineCount(fileBytes)
	wordCount = ccwc.FindWordCount(fileBytes)
	return 
}

func openFile(filename string) io.Reader {
	reader, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return reader
}

// func displayHelpInfo() {
// 	fmt.Println("Usage:")
// 	fmt.Println("ccwc [OPTION] [FILE]")
// 	fmt.Println("Print newline, word, character and byte counts for the specified FILE.")
// 	fmt.Print("A word is a non-zero-length sequence of characters delimited by white space.\n\n")
// 	fmt.Println("The options below may be used to select which counts are printed.")
// 	fmt.Println("-c               print the byte count")
// 	fmt.Println("-m               print the character count")
// 	fmt.Println("-w               print the word count")
// 	fmt.Println("-l               print the line count")
// 	fmt.Println("--help           display this help and exit")
// 	fmt.Print("\nIf no optional flags is entered, it prints the file information for all flags except the help flag\n")
// 	fmt.Println("Full documentation <https://github.com/emanuelquerty/ccwc>")
// }
