package tool

import "cxor/tool/dummy"

type Tool interface {
	Encrypt(data string, key string) (encryptedData string, err error)
	Decrypt(encryptedData string, key string) (decryptedData string, err error)
}

func NewTool() Tool {
	return &dummy.ToolV1{}
}
