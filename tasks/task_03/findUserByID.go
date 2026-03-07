package task03

import "errors"

type User struct {
	ID int
	Name string
	Age int
}

func FindUserByID(users []User, id int) (*User, error) {
	//your code will be here
	return nil, errors.New("non implemented") //заглушка
}