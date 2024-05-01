package ccwc

import "bytes"

func FindByteCount(fileBytes []byte) int {
	return len(fileBytes)
}

func FindLineCount(fileBytes []byte) int {
	return bytes.Count(fileBytes, []byte{'\n'})
}

func FindWordCount(fileBytes []byte) int {
	return len(bytes.Fields(fileBytes))
}

func FindCharacterCount(fileBytes []byte) int {
	return len(bytes.Runes(fileBytes))
}