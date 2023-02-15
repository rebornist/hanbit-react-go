package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"hanbit-react/util/component"
)

type Crypt struct {
	cipher.Block
}

func NewCryptService() (CryptService, error) {

	conf, err := component.NewJSONConfig("config_local.json")
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

	return &Crypt{cipher}, nil
}
