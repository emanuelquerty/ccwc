package ccwc_test

import (
	"log"
	"os"
	"testing"

	"github.com/emanuelquerty/ccwc"
	"github.com/emanuelquerty/ccwc/testdata"
	"github.com/stretchr/testify/assert"
)


func TestFindByteCount(t *testing.T) {
	errMessage := "Byte Count: the two numbers should be the same"
	assert := assert.New(t)

	for _, test := range testdata.TableDrivenData {
		fileBytes := readFile(t, test.Filepath)
		got := ccwc.FindByteCount(fileBytes)
		assert.Equal(test.Expected.Bytes, got, errMessage)
	}
}

func TestFindLineCount(t *testing.T) {
	errMessage := "Line Count: the two numbers should be the same"
	assert := assert.New(t)

	for _, test := range testdata.TableDrivenData {
		fileBytes := readFile(t, test.Filepath)
		got := ccwc.FindLineCount(fileBytes)
		assert.Equal(test.Expected.Lines, got, errMessage)
	}
}

func TestFindWordCount(t *testing.T) {
	errMessage := "Word Count: the two numbers should be the same"
	assert := assert.New(t)

	
	for _, test := range testdata.TableDrivenData {
		fileBytes := readFile(t, test.Filepath)
		got := ccwc.FindWordCount(fileBytes)
		assert.Equal(test.Expected.Words, got, errMessage)
	}
}

func TestFindCharacterCount(t *testing.T) {
	errMessage := "Chacter Count: the two numbers should be the same"
	assert := assert.New(t)
	
	for _, test := range testdata.TableDrivenData {
		fileBytes := readFile(t, test.Filepath)
		got := ccwc.FindCharacterCount(fileBytes)
		assert.Equal(test.Expected.Chars, got, errMessage)
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
