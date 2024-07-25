package dummy

type ToolV1 struct{}

func (t *ToolV1) Encrypt(plaintext string, key string) (string, error) {
	return XorEncryptDecrypt(Encrypt(plaintext, len(key)), key), nil
}

func (t *ToolV1) Decrypt(ciphertext string, key string) (string, error) {
	return Decrypt(XorEncryptDecrypt(ciphertext, key), len(key)), nil
}
