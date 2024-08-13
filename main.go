package main

import (
	"cxor/tool"
	"fmt"
)

func main() {
	data := "Hello, World! 123 *xd"
	key := "123Key 1"

	tool := tool.NewTool()
	fmt.Println("Tool:", tool)

	encrypted, err := tool.Encrypt(data, key)
	fmt.Printf("err: %v\n", err)
	fmt.Println("Encrypted:", encrypted)

	decrypted, err := tool.Decrypt(encrypted, key)
	fmt.Printf("err: %v\n", err)
	fmt.Println("Decrypted:", decrypted)

	fmt.Println(">", data == decrypted)
}

// func EncryptCaesarCipherXOR(plaintext string, key string) string {
// 	return xorEncryptDecrypt(encrypt(plaintext, len(key)), key)
// }

// func DecryptCaesarCipherXOR(ciphertext string, key string) string {
// 	return decrypt(xorEncryptDecrypt(ciphertext, key), len(key))
// }

// func xorEncryptDecrypt(data string, key string) string {
// 	encrypted := make([]byte, len(data))
// 	keyLength := len(key)
// 	for i := 0; i < len(data); i++ {
// 		encrypted[i] = data[i] ^ key[i%keyLength]
// 	}
// 	return string(encrypted)
// }

// // encrypt encrypts the given plaintext using the Caesar cipher with the specified shift
// func encrypt(plaintext string, shift int) string {
// 	var encrypted strings.Builder

// 	for _, char := range plaintext {

// 		if char >= 'a' && char <= 'z' {
// 			char = rune((int(char)-'a'+shift)%26 + 'a')
// 		} else if char >= 'A' && char <= 'Z' {
// 			char = rune((int(char)-'A'+shift)%26 + 'A')
// 		} else if char >= '0' && char <= '9' {
// 			char = rune((int(char)-'0'+(shift+2))%10 + '0')
// 		}

// 		encrypted.WriteByte(byte(char))
// 	}

// 	return encrypted.String()
// }

// // decrypt decrypts the given ciphertext using the Caesar cipher with the specified shift
// func decrypt(ciphertext string, shift int) string {
// 	return encrypt(ciphertext, 26-shift)
// }
