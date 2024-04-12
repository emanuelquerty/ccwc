package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestFindByteCount(t *testing.T) {
	errMessage := "Byte Count: the two numbers should be the same"
	assert := assert.New(t)

	for _, test := range testData {
		fileBytes := readFile(t, test.filepath)
		got := findByteCount(fileBytes)
		assert.Equal(test.expected.bytes, got, errMessage)
	}
}

func TestFindLineCount(t *testing.T) {
	errMessage := "Line Count: the two numbers should be the same"
	assert := assert.New(t)

	for _, test := range testData {
		fileBytes := readFile(t, test.filepath)
		got := findLineCount(fileBytes)
		assert.Equal(test.expected.lines, got, errMessage)
	}
}

func TestFindWordCount(t *testing.T) {
	errMessage := "Word Count: the two numbers should be the same"
	assert := assert.New(t)

	
	for _, test := range testData {
		fileBytes := readFile(t, test.filepath)
		got := findWordCount(fileBytes)
		assert.Equal(test.expected.words, got, errMessage)
	}
}

func TestFindCharacterCount(t *testing.T) {
	errMessage := "Chacter Count: the two numbers should be the same"
	assert := assert.New(t)
	
	for _, test := range testData {
		fileBytes := readFile(t, test.filepath)
		got := findCharacterCount(fileBytes)
		assert.Equal(test.expected.chars, got, errMessage)
	}
}

func readFile(t testing.TB, filepath string) []byte {
	t.Helper()
	fileBytes, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return fileBytes
}
