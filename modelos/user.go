package modelos

import (
	"time"
)

type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdateAt  time.Time
	Status    bool
}

func (u *User) AddUser(id int, name string) {
	u.Id = id
	u.Name = name
	u.CreatedAt = time.Now()
	u.UpdateAt = time.Now()
	u.Status = true
}
