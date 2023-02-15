package auth

import "hanbit-react/types"

type AuthService interface {
	BasicAuthentication(string, string) (types.User, error)
}
