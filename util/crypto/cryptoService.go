package crypto

type CryptoService interface {
	EncryptAES(string) (string, error)
	DecryptAES(string) (string, error)
}
