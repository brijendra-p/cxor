package v1

func XorEncryptDecrypt(data string, key string) string {
	encrypted := make([]byte, len(data))
	keyLength := len(key)
	for i := 0; i < len(data); i++ {
		encrypted[i] = data[i] ^ key[i%keyLength]
	}
	return string(encrypted)
}
