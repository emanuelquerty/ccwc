package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var fileBytes1, _ = os.ReadFile("test.txt")
var fileBytes2, _ = os.ReadFile("test2.txt")
var fileBytes3, _ = os.ReadFile("test3.txt")
var fileBytes4, _ = os.ReadFile("test4.txt")

func TestFindByteCount(t *testing.T) {
	actual1 := findByteCount(fileBytes1)
	actual2 := findByteCount(fileBytes2)
	actual3 := findByteCount(fileBytes3)
	actual4 := findByteCount(fileBytes4)

	assert := assert.New(t)
	assert.Equal(342190, actual1, "Byte Count: the two numbers should be the same")
	assert.Equal(2, actual2, "Byte Count: the two numbers should be the same")
	assert.Equal(738, actual3, "Byte Count: the two numbers should be the same")
	assert.Equal(44, actual4, "Byte Count: the two numbers should be the same")
}

func TestFindLineCount(t *testing.T) {
	actual1 := findLineCount(fileBytes1)
	actual2 := findLineCount(fileBytes2)
	actual3 := findLineCount(fileBytes3)
	actual4 := findLineCount(fileBytes4)

	assert := assert.New(t)
	assert.Equal(7145, actual1, "Line Count: the two numbers should be the same")
	assert.Equal(1, actual2, "Line Count: the two numbers should be the same")
	assert.Equal(9, actual3, "Line Count: the two numbers should be the same")
	assert.Equal(5, actual4, "Line Count: the two numbers should be the same")
}

func TestFindWordCount(t *testing.T) {
	actual1 := findWordCount(fileBytes1)
	actual2 := findWordCount(fileBytes2)
	actual3 := findWordCount(fileBytes3)
	actual4 := findWordCount(fileBytes4)

	assert := assert.New(t)
	assert.Equal(58164, actual1, "Word Count: the two numbers should be the same")
	assert.Equal(0, actual2, "Word Count: the two numbers should be the same")
	assert.Equal(113, actual3, "Word Count: the two numbers should be the same")
	assert.Equal(6, actual4, "Word Count: the two numbers should be the same")
}

func TestFindCharacterCount(t *testing.T) {
	actual1 := findCharacterCount(fileBytes1)
	actual2 := findCharacterCount(fileBytes2)
	actual3 := findCharacterCount(fileBytes3)
	actual4 := findCharacterCount(fileBytes4)

	assert := assert.New(t)
	assert.Equal(339292, actual1, "Character Count: the two numbers should be the same")
	assert.Equal(2, actual2, "Character Count: the two numbers should be the same")
	assert.Equal(738, actual3, "Character Count: the two numbers should be the same")
	assert.Equal(44, actual4, "Character Count: the two numbers should be the same")
}
