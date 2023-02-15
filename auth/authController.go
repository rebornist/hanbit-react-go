package auth

import (
	"fmt"
	"hanbit-react/types"
	"hanbit-react/util/crypt"
)

// Basic Authentication
func (a *Auth) BasicAuthentication(username string, password string) (types.User, error) {

	// 유저 정보 가져오기
	user := types.User{}

	// 유저 정보 가져오기
	user, err := a.db.GetUserByUsername(username)
	if err != nil {
		return types.User{}, err
	}

	// 유저 정보가 없으면 false
	if user.Username == "" {
		return types.User{}, fmt.Errorf("not found user")
	}

	// 비밀번호 비교 기능
	// 암호화 기능 선언
	crypt, err := crypt.NewCryptService()
	if err != nil {
		return types.User{}, err
	}

	// 패스워드 복호화
	decryptedPassword, err := crypt.DecryptAES(user.Password)
	if err != nil {
		return types.User{}, err
	}

	// 복호화된 패스워드와 입력된 패스워드 비교
	if decryptedPassword != password {
		return types.User{}, nil
	}

	return user, nil
}
