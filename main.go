package main

import (
	"strings"
)

func caesarCipher(text string, number int) string {
	var result strings.Builder
	number %= 26
	for _, char := range text {
		switch {
		case 'a' <= char && char <= 'z':
			result.WriteByte(byte('a' + (char-'a'+rune(number))%26))
		case 'A' <= char && char <= 'Z':
			result.WriteByte(byte('A' + (char-'A'+rune(number))%26))
		default:
			result.WriteByte(byte(char))
		}
	}
	return result.String()
}

func main() {

}
