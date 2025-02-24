package user

import "fmt"

type user struct {
	name string
	age  int
}

func New(name string, age int) *user {
	return &user{name, age}
}
func (s *user) start() string {
	return fmt.Sprint("starting")
}
