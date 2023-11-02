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

	flag.Parse()

	if flagvar == "" {
		flagvar = flag.Arg(0)
		if flagvar == "" {
			fmt.Print("No path to a file was specified.\n\n")
			fmt.Println("Please, enter a file location in the following format:")
			fmt.Print("./ccwc [optional flags] [filepath]\n\n")
			fmt.Println("Available optional flags:")
			fmt.Println("-c : Counts the number of bytes in the file")
			fmt.Println("-l : Counts the number of lines in the file")
			fmt.Println("-w : Counts the number of words in the file")
			fmt.Println("-m : Counts the number of characters in the file")
			fmt.Print("\nIf no optional flags is entered, it prints the file information for all flags\n")
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
		fmt.Printf(" %d %s", byteCount, flagvar)
	case "l":
		lineCount = findLineCount(fileBytes)
		fmt.Printf(" %d %s", lineCount, flagvar)
	case "w":
		wordCount = findWordCount(fileBytes)
		fmt.Printf(" %d %s", wordCount, flagvar)
	case "m":
		charCount = findCharacterCount(fileBytes)
		fmt.Printf(" %d %s", charCount, flagvar)
	default:
		byteCount, lineCount, wordCount = findAll(fileBytes)
		fmt.Printf(" %d %d %d %s", lineCount, wordCount, byteCount, flagvar)
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
