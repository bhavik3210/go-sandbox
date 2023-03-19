package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextId = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New User must not include id or it must be set to 0")
	}
	u.ID = nextId
	nextId++
	users = append(users, &u)
	return u, nil
}

// if errors are non-nill DO NOT ignore them,
// handle them properly because it is an indicator something really bad has happened.
func GetUserById(id int) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return *user, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

func UpdateUser(user User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == user.ID {
			users[i] = &user
			return user, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", user.ID)
}

func RemoveUserById(id int) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}
