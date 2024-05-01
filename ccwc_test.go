package ccwc_test

import (
	"bytes"
	"os"
	"reflect"
	"testing"

	"github.com/emanuelquerty/ccwc"
	"github.com/emanuelquerty/ccwc/testdata"
)

func TestParseFlag(t *testing.T) {
	t.Run("no flags have been set", func(t *testing.T) {
		flag, ok := ccwc.GetFlag()

		if ok {
			t.Errorf("expected no flags set but got %+v", flag)
		}
	})

	t.Run("defines a flag", func(t *testing.T) {
		cflag := ccwc.Flag {
			Name: "c",
			Usage: "outputs the byte counts",
		}

		isDefined := ccwc.DefineFlag(cflag.Name, cflag.Usage)
		if !isDefined {
			t.Error("expected flag to be defined but flag does not exist")
		}
	})

	t.Run("sets a flag to a given value", func(t *testing.T) {
		ccwc.ResetFlagSet()
		mFlag := ccwc.Flag {
			Name: "m",
			Value: "post1.md",
			Usage: "outputs the character counts",
		}

		args := []string{"-m", "post1.md"}

		_ = ccwc.DefineFlag(mFlag.Name, mFlag.Usage)
		err := ccwc.SetFlag(args)

		if err != nil {
			t.Error(err)
		}

		actualFlag, ok := ccwc.GetFlag()
		if !ok {
			t.Errorf("expected one flag but got %+v", actualFlag)
		}

		if !reflect.DeepEqual(mFlag, actualFlag) {
			t.Errorf("expected %+v but got %+v", mFlag, actualFlag)
		}
	})

	t.Run("sets a flag but passes no value", func(t *testing.T) {
		ccwc.ResetFlagSet()
		lFlaG := ccwc.Flag {
			Name: "l",
			Value: "",
			Usage: "outputs the character counts",
		}
		args := []string{"-l"}
		_= ccwc.DefineFlag(lFlaG.Name, lFlaG.Usage)
		err := ccwc.SetFlag(args)
		
		if err != nil {
			t.Error(err)
		}

		actualFlag, ok := ccwc.GetFlag()
		if !ok {
			t.Errorf("expected one flag but got %+v", actualFlag)
		}

		if !reflect.DeepEqual(lFlaG, actualFlag) {
			t.Errorf("expected %+v but got %+v", lFlaG, actualFlag)
		}

	})

	t.Run("resets the flag by creating a new flagset", func(t *testing.T) {
		lFlaG := ccwc.Flag {
			Name: "w",
			Value: "post1.md",
			Usage: "outputs the character counts",
		}
		args := []string{"-w", "post1.md"}
		_= ccwc.DefineFlag(lFlaG.Name, lFlaG.Usage)
		_= ccwc.SetFlag(args)
		
		ccwc.ResetFlagSet()
		actualFlag, ok := ccwc.GetFlag()
		if ok {
			t.Errorf("expected no flag but got flag %+v", actualFlag)
		}
	})
}

func TestReadFile(t *testing.T) {
	fileContent := []byte("DummyText\n SomeMoreText\n")
	buff := bytes.NewReader(fileContent)

	got, err := ccwc.ReadFile(buff)
	if err != nil {
		t.Fatal(err)
	}

	if string(fileContent) != string(got) {
		t.Fatalf("expected %v but got %v", string(fileContent), string(got))
	}
}

func TestCounts(t *testing.T) {	
	for _, rowData := range testdata.TableDrivenData {
		file, err := os.Open(rowData.Filepath)
		if err != nil {
			t.Fatal(err)
		}

		fileContent, _ := ccwc.ReadFile(file)

		gotChars := ccwc.FindCharacterCount(fileContent)
		gotBytes := ccwc.FindByteCount(fileContent)
		gotWords := ccwc.FindWordCount(fileContent)
		gotLines := ccwc.FindLineCount(fileContent)

		assertHelper(t, rowData.Expected.Chars, gotChars, "expected %d chars but got %d chars")
		assertHelper(t, rowData.Expected.Bytes, gotBytes, "expected %d bytes but got %d bytes")
		assertHelper(t, rowData.Expected.Words, gotWords, "expected %d words but got %d words")
		assertHelper(t, rowData.Expected.Lines, gotLines, "expected %d lines but got %d lines")
	}
}

func assertHelper(t *testing.T, expected int, got int, errMessage string) {
	t.Helper()
	if expected != got {
		t.Errorf(errMessage, expected, got)
	}
}