package entity

import "fmt"

type User struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

func NewUser(id string, name string) (user User) {
	user = User{ID: id}
	user.SetName(name)
	return user
}

func (p *User) SetName(name string) {
	p.Name = fmt.Sprintf("%.10s", name)
}
