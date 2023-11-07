package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var fileBytes1, _ = os.ReadFile("testfiles/test.txt")
var fileBytes2, _ = os.ReadFile("testfiles/test2.txt")
var fileBytes3, _ = os.ReadFile("testfiles/test3.txt")
var fileBytes4, _ = os.ReadFile("testfiles/test4.txt")
var fileBytes5, _ = os.ReadFile("testfiles/chineseRandom.txt")

// var fileBytes6, _ = os.ReadFile("testfiles/hugeRandom1GB.txt")
var fileBytes7, _ = os.ReadFile("testfiles/oneLongWord.txt")

var byteActual1, lineActual1, wordActual1 = findAll(fileBytes1)
var byteActual2, lineActual2, wordActual2 = findAll(fileBytes2)
var byteActual3, lineActual3, wordActual3 = findAll(fileBytes3)
var byteActual4, lineActual4, wordActual4 = findAll(fileBytes4)
var byteActual5, lineActual5, wordActual5 = findAll(fileBytes5)

// var byteActual6, lineActual6, wordActual6 = findAll(fileBytes6)
var byteActual7, lineActual7, wordActual7 = findAll(fileBytes7)

func TestFindByteCount(t *testing.T) {
	errMessage := "Byte Count: the two numbers should be the same"
	assert := assert.New(t)

	assert.Equal(342190, byteActual1, errMessage)
	assert.Equal(2, byteActual2, errMessage)
	assert.Equal(738, byteActual3, errMessage)
	assert.Equal(44, byteActual4, errMessage)
	assert.Equal(1613, byteActual5, errMessage)
	// assert.Equal(1073741824, byteActual6, errMessage)
	assert.Equal(137, byteActual7, errMessage)

}

func TestFindLineCount(t *testing.T) {
	errMessage := "Line Count: the two numbers should be the same"
	assert := assert.New(t)

	assert.Equal(7145, lineActual1, errMessage)
	assert.Equal(1, lineActual2, errMessage)
	assert.Equal(9, lineActual3, errMessage)
	assert.Equal(5, lineActual4, errMessage)
	assert.Equal(13, lineActual5, errMessage)
	// assert.Equal(13944699, lineActual6, errMessage)
	assert.Equal(0, lineActual7, errMessage)
}

func TestFindWordCount(t *testing.T) {
	errMessage := "Word Count: the two numbers should be the same"
	assert := assert.New(t)

	assert.Equal(58164, wordActual1, errMessage)
	assert.Equal(0, wordActual2, errMessage)
	assert.Equal(113, wordActual3, errMessage)
	assert.Equal(6, wordActual4, errMessage)
	assert.Equal(12, wordActual5, errMessage)
	// assert.Equal(13944700, wordActual6, errMessage)
	assert.Equal(1, wordActual7, errMessage)
}

func TestFindCharacterCount(t *testing.T) {
	actual1 := findCharacterCount(fileBytes1)
	actual2 := findCharacterCount(fileBytes2)
	actual3 := findCharacterCount(fileBytes3)
	actual4 := findCharacterCount(fileBytes4)

	actual5 := findCharacterCount(fileBytes5)
	// actual6 := findCharacterCount(fileBytes6)
	actual7 := findCharacterCount(fileBytes7)

	errMessage := "Chacter Count: the two numbers should be the same"
	assert := assert.New(t)
	assert.Equal(339292, actual1, errMessage)
	assert.Equal(2, actual2, errMessage)
	assert.Equal(738, actual3, errMessage)
	assert.Equal(44, actual4, errMessage)
	assert.Equal(561, actual5, errMessage)
	// assert.Equal(1073741824, actual6, errMessage)
	assert.Equal(137, actual7, errMessage)
}
