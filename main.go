package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func caesarCipher(text string, shift int) string {
	shift %= 26
	var result strings.Builder
	result.Grow(len(text))

	for _, char := range text {
		if 'a' <= char && char <= 'z' {
			result.WriteByte(byte('a' + (char-'a'+rune(shift))%26))
		} else if 'A' <= char && char <= 'Z' {
			result.WriteByte(byte('A' + (char-'A'+rune(shift))%26))
		} else {
			result.WriteByte(byte(char))
		}
	}
	return result.String()
}

func main() {
	fmt.Print("Enter password: ")
	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Failed to read password:", err)
		return
	}
	fmt.Println("\nPassword accepted.")

	password := string(bytePassword)
	expectedPassword := "secret123" // Здесь ваш секретный пароль

	if password != expectedPassword {
		fmt.Println("Invalid password. Access denied.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter file path: ")
	filePath, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read file path:", err)
		return
	}
	filePath = strings.TrimSpace(filePath)

	fmt.Print("Choose mode (encrypt/decrypt): ")
	mode, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read mode:", err)
		return
	}
	mode = strings.TrimSpace(mode)

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}
	text := string(fileContent)

	shift := 3 // Фиксированный сдвиг, можно запросить у пользователя или делать его частью пароля
	if mode == "decrypt" {
		shift = -shift
	}

	result := caesarCipher(text, shift)
	fmt.Println("Result:")
	fmt.Println(result)
}
