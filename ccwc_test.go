package ccwc_test

import (
	"log"
	"os"
	"testing"

	"github.com/emanuelquerty/ccwc"
	"github.com/emanuelquerty/ccwc/testdata"
)


func TestFindByteCount(t *testing.T) {
	errMessage := "Byte Count: the two numbers should be the same"

	for _, test := range testdata.TableDrivenData {
		fileBytes := readFile(t, test.Filepath)
		got := ccwc.FindByteCount(fileBytes)
		assertHelper(t, test.Expected.Bytes, got, errMessage)
	}
}

func TestFindLineCount(t *testing.T) {
	errMessage := "Line Count: the two numbers should be the same"

	for _, test := range testdata.TableDrivenData {
		fileBytes := readFile(t, test.Filepath)
		got := ccwc.FindLineCount(fileBytes)
		assertHelper(t, test.Expected.Lines, got, errMessage)
	}
}

func TestFindWordCount(t *testing.T) {
	errMessage := "Word Count: the two numbers should be the same"

	
	for _, test := range testdata.TableDrivenData {
		fileBytes := readFile(t, test.Filepath)
		got := ccwc.FindWordCount(fileBytes)
		assertHelper(t, test.Expected.Words, got, errMessage)
	}
}

func TestFindCharacterCount(t *testing.T) {
	errMessage := "Chacter Count: the two numbers should be the same"
	
	for _, test := range testdata.TableDrivenData {
		fileBytes := readFile(t, test.Filepath)
		got := ccwc.FindCharacterCount(fileBytes)
		assertHelper(t, test.Expected.Chars, got, errMessage)
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

func assertHelper(t *testing.T, expected int, got int, errMessage string) {
	t.Helper()
	if expected != got {
		t.Errorf(errMessage, expected, got)
	}
}