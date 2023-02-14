package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"hanbit-react/util/component"
)

type Crypto struct {
	cipher.Block
}

func NewCryptoService() (CryptoService, error) {

	conf, err := component.NewJSONConfig("config.json")
	if err != nil {
		return nil, err
	}

	// config 정보 받아오기
	byteData, err := conf.GetServiceInfo("private_key")
	if err != nil {
		return nil, err
	}

	cipher, err := aes.NewCipher(byteData)
	if err != nil {
		return nil, err
	}

	return &Crypto{cipher}, nil
}
