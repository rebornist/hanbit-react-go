package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base32"
	"fmt"
)

func (c *Crypt) DecryptAES(txt string) (string, error) {

	// 암호화된 문자열 byte 타입으로 변환
	ciphertext, err := base32.StdEncoding.DecodeString(txt)
	if err != nil {
		return "", err
	} else if len(ciphertext)%aes.BlockSize != 0 { // 블록 크기의 배수가 아니면 리턴
		return "", fmt.Errorf("%s", "암호화된 데이터의 길이는 블록 크기의 배수가 되어야합니다.")
	}

	iv := ciphertext[:aes.BlockSize]        // 부분 슬라이스로 초기화 벡터 공간을 가져옴
	ciphertext = ciphertext[aes.BlockSize:] // 부분 슬라이스로 암호화된 데이터를 가져옴

	plaintext := make([]byte, len(ciphertext))  // 평문 데이터를 저장할 공간 생성
	mode := cipher.NewCBCDecrypter(c.Block, iv) // 암호화 블록과 초기화 벡터를 넣어서

	// 복호화 블록 모드 인스턴스 생성
	mode.CryptBlocks(plaintext, ciphertext) // 복호화 블록 모드 인스턴스로 복호화

	// 빈 문자열 제거
	plaintext = bytes.Trim(plaintext, "\x00")

	// 문자열로 변환
	unsignedtext := string(plaintext)

	return unsignedtext, nil

}
