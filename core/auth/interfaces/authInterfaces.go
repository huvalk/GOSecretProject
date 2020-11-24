package authInterfaces

import (
	"GOSecretProject/core/model/base"
)

type AuthRepository interface {
	Register(user baseModels.User) (err error)
	Login(user baseModels.User) (userID int, session string, statusCode int, err error)
	Logout(session string) (err error)
	CheckSession(session string) (userId uint64, err error)
}
