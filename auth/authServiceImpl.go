package auth

import "hanbit-react/mapper/user"

type Auth struct {
	db user.UserSerivce
}

func NewAuthService() (AuthService, error) {
	// 유저 서비스 선언
	user, err := user.NewUserService()
	if err != nil {
		return nil, err
	}

	return &Auth{db: user}, nil
}
