package User

import (
	"mock/IUser"
)

type User struct {
	IUser IUser.Mako
}

func (u *User) Use() {
	u.IUser.AddUser(1, "sample mock test")
}
