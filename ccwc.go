/*
Package ccwc implements required methods for word count (ccwc)
*/

package ccwc

import (
	"bytes"
	"flag"
	"io"
	"strings"
)


var flagSet = NewFlagSet()
var flagValue = FlagValue{}

// Creates a new flagset
func NewFlagSet() *flag.FlagSet{
	return flag.NewFlagSet("ccwc", flag.ExitOnError)
}

// Returns the currently defined flagset
func FlagSet() *flag.FlagSet {
	return flagSet
}

// Resets the flagset
func ResetFlagSet() {
	flagSet = NewFlagSet()
}

// Gets the flag that has been set
func TargetFlag() (Flag, bool) {
	var hasFlags bool
	var foundFlag Flag

	flagSet.Visit(func(f *flag.Flag) {
		foundFlag = Flag{
			Name: f.Name,
			Value: f.Value.String(),
			Usage: f.Usage,
		}
		hasFlags = true
	})
	return foundFlag, hasFlags
}

// Defines a flag but does not set it to any value
func DefineFlag(name string, usage string) bool {
	flagSet.Var(&flagValue, name, usage)
	
	var isDefined bool
	flagSet.VisitAll(func(f *flag.Flag) {
		isDefined = true
	})
	return isDefined
}

// A flag is any character or word preceeded by "-" or "--"
// Sets sets a flag to its given value. 
// A flag passed in with no value gets a default value of ""
func SetFlag(args []string) error {
	newArgs := make([]string, len(args))
	copy(newArgs, args)
	if len(args) == 1 && strings.HasPrefix(args[0], "-") {
		newArgs = append(newArgs, "" )
	}

	err := flagSet.Parse(newArgs)
	return err
}

// Reads the content of a reader and returns the bytes read
func ReadFile(reader io.Reader) ([]byte, error) {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return []byte(""), err
	}
	return bytes, nil
}

func FindCharacterCount(fileBytes []byte) int {
	return len(bytes.Runes(fileBytes))
}

func FindByteCount(fileBytes []byte) int {
	return len(fileBytes)
}

func FindLineCount(fileBytes []byte) int {
	return bytes.Count(fileBytes, []byte{'\n'})
}

func FindWordCount(fileBytes []byte) int {
	return len(bytes.Fields(fileBytes))
}

