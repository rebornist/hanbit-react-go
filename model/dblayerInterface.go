package model

import (
	"hanbit-react/types"
)

type DBLayerInterface interface {
	GetAllUsers() ([]types.User, error)
}
