package tool

import v1 "cxor/tool/v1"

type Tool interface {
	Encrypt(data string, key string) (encryptedData string, err error)
	Decrypt(encryptedData string, key string) (decryptedData string, err error)
}

func NewTool() Tool {
	tool := &v1.ToolV1{}
	tool.Token.SetUnlimited()
	return tool
}
