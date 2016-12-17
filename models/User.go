package models

import (
	"github.com/revel/revel"
	"fmt"
)

type User struct {
	Name string
}

func (u *User) String() string {
	return fmt.Sprintf("User(%s)", u.Name)
}

func (user *User) Validate(v *revel.Validation) {
	v.Required(user.Name).Message("A valid user name is requried")
	v.MinSize(user.Name, 4).Message("Min 4")
	v.MaxSize(user.Name, 15).Message("Max 15")
}