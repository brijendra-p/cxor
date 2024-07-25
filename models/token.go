package models

type Token struct {
	Encrypt   int
	Decrypt   int
	Unlimited bool
}

func (t *Token) SetEncryptToken(token int) {
	t.Encrypt = token
}

func (t *Token) GetEncryptToken() int {
	return t.Encrypt
}

func (t *Token) UsingEncryptToken() {
	t.Encrypt -= 1
}

func (t *Token) SetDecryptToken(token int) {
	t.Decrypt = token
}

func (t *Token) GetDecryptToken() int {
	return t.Decrypt
}

func (t *Token) UsingDecryptToken() {
	t.Decrypt -= 1
}

func (t *Token) SetUnlimited() {
	t.Unlimited = true
}
func (t *Token) IsUnlimited() bool {
	return t.Unlimited
}
