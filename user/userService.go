package user

import (
	"hanbit-react/types"
)

type UserSerivce interface {
	GetAllUsers() ([]types.User, error)
}
