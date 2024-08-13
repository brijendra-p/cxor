package models

type Token struct {
	encrypt   int
	decrypt   int
	unlimited bool
}

func (t *Token) SetEncryptToken(token int) {
	t.encrypt = token
}

func (t *Token) GetEncryptToken() int {
	return t.encrypt
}

func (t *Token) UsingEncryptToken() {
	t.encrypt -= 1
}

func (t *Token) SetDecryptToken(token int) {
	t.decrypt = token
}

func (t *Token) GetDecryptToken() int {
	return t.decrypt
}

func (t *Token) UsingDecryptToken() {
	t.decrypt -= 1
}

func (t *Token) SetUnlimited() {
	t.unlimited = true
}
func (t *Token) IsUnlimited() bool {
	return t.unlimited
}
