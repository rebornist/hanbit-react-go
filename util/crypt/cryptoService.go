package crypt

type CryptService interface {
	EncryptAES(string) (string, error)
	DecryptAES(string) (string, error)
}
