package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var flagvar string
	flag.StringVar(&flagvar, "c", "", "Counts the number of bytes in the file specified by the flag")
	flag.StringVar(&flagvar, "l", "", "Counts the number of lines in the file specified by the flag")
	flag.StringVar(&flagvar, "w", "", "Counts the number of words in the file specified by the flag")
	flag.StringVar(&flagvar, "m", "", "Counts the number of characters in the file specified by the flag")

	flag.Usage = displayHelpInfo
	flag.Parse()

	if flagvar == "" {
		flagvar = flag.Arg(0)
		if flagvar == "" {
			fmt.Print("No path to a file was specified.\n\n")
			flag.Usage()
			return
		}
	}

	fileBytes, err := os.ReadFile(flagvar)
	if err != nil {
		log.Fatalf("ccwc: %s: No such file or directory", flagvar)
	}

	flagName := findWhichFlag()
	var byteCount, lineCount, wordCount, charCount int
	switch flagName {
	case "c":
		byteCount = findByteCount(fileBytes)
		fmt.Printf(" %d %s\n", byteCount, flagvar)
	case "l":
		lineCount = findLineCount(fileBytes)
		fmt.Printf(" %d %s\n", lineCount, flagvar)
	case "w":
		wordCount = findWordCount(fileBytes)
		fmt.Printf(" %d %s\n", wordCount, flagvar)
	case "m":
		charCount = findCharacterCount(fileBytes)
		fmt.Printf(" %d %s\n", charCount, flagvar)
	default:
		byteCount, lineCount, wordCount = findAll(fileBytes)
		fmt.Printf(" %d %d %d %s\n", lineCount, wordCount, byteCount, flagvar)
	}
}

func findWhichFlag() string {
	var flagName string
	flag.Visit(func(f *flag.Flag) {
		flagName = f.Name
	})
	return flagName
}

func findAll(fileBytes []byte) (int, int, int) {
	byteCount := findByteCount(fileBytes)
	lineCount := findLineCount(fileBytes)
	wordCount := findWordCount(fileBytes)
	return byteCount, lineCount, wordCount
}

func findByteCount(fileBytes []byte) int {
	return len(fileBytes)
}

func findLineCount(fileBytes []byte) int {
	return bytes.Count(fileBytes, []byte{'\n'})
}

func findWordCount(fileBytes []byte) int {
	return len(bytes.Fields(fileBytes))
}

func findCharacterCount(fileBytes []byte) int {
	return len(bytes.Runes(fileBytes))
}

func displayHelpInfo() {
	fmt.Println("Usage:")
	fmt.Println("ccwc [OPTION] [FILE]")
	fmt.Println("Print newline, word, character and byte counts for the specified FILE.")
	fmt.Print("A word is a non-zero-length sequence of characters delimited by white space.\n\n")
	fmt.Println("The options below may be used to select which counts are printed.")
	fmt.Println("-c               print the byte count")
	fmt.Println("-m               print the character count")
	fmt.Println("-w               print the word count")
	fmt.Println("-l               print the line count")
	fmt.Println("--help           display this help and exit")
	fmt.Print("\nIf no optional flags is entered, it prints the file information for all flags except the help flag\n")
	fmt.Println("Full documentation <https://github.com/emanuelquerty/ccwc>")
}
