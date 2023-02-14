package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base32"
	"io"
)

func (c *Crypto) EncryptAES(txt string) (string, error) {

	// 전달받은 문자열 byte 타입으로 변환
	plaintext := []byte(txt)
	if mod := len(plaintext) % aes.BlockSize; mod != 0 { // 블록 크기의 배수가 되어야함
		padding := make([]byte, aes.BlockSize-mod) // 블록 크기에서 모자라는 부분을
		plaintext = append(plaintext, padding...)  // 채워줌
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext)) // 초기화 벡터 공간(aes.BlockSize)만큼 더 생성
	iv := ciphertext[:aes.BlockSize]                         // 부분 슬라이스로 초기화 벡터 공간을 가져옴
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {  // 랜덤 값을 초기화 벡터에 넣어줌
		return "", err
	}

	mode := cipher.NewCBCEncrypter(c.Block, iv)             // 암호화 블록과 초기화 벡터를 넣어서 암호화 블록 모드 인스턴스 생성
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext) // 암호화 블록 모드 인스턴스로 암호화

	// 문자열로 변환
	signedtext := base32.StdEncoding.EncodeToString(ciphertext)

	return signedtext, nil

}
