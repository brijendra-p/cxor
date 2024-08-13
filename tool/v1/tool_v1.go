package v1

import (
	"cxor/models"
	"fmt"
)

type ToolV1 struct {
	Token models.Token
}

func (t *ToolV1) Encrypt(plaintext string, key string) (string, error) {
	if !t.Token.IsUnlimited() && t.Token.GetEncryptToken() == 0 {
		return "", fmt.Errorf("not have enough tokens for encrypt")
	}

	return XorEncryptDecrypt(Encrypt(plaintext, len(key)), key), nil
}

func (t *ToolV1) Decrypt(ciphertext string, key string) (string, error) {
	if !t.Token.IsUnlimited() && t.Token.GetEncryptToken() == 0 {
		return "", fmt.Errorf("not have enough tokens for decrypt")
	}
	return Decrypt(XorEncryptDecrypt(ciphertext, key), len(key)), nil
}
