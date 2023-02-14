package auth

import (
	"hanbit-react/types"
)

type AuthSerivce interface {
	GetAllUsers() ([]types.User, error)
}
