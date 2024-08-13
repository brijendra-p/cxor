package v1

import "strings"

// Encrypt encrypts the given plaintext using the Caesar cipher with the specified shift
func Encrypt(plaintext string, shift int) string {
	var encrypted strings.Builder

	for _, char := range plaintext {

		if char >= 'a' && char <= 'z' {
			char = rune((int(char)-'a'+shift)%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			char = rune((int(char)-'A'+shift)%26 + 'A')
		} else if char >= '0' && char <= '9' {
			char = rune((int(char)-'0'+(shift+2))%10 + '0')
		}

		encrypted.WriteByte(byte(char))
	}

	return encrypted.String()
}

// Decrypt decrypts the given ciphertext using the Caesar cipher with the specified shift
func Decrypt(ciphertext string, shift int) string {
	return Encrypt(ciphertext, 26-shift)
}
