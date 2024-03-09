package main

import (
	"errors"
)

type AuthService struct {
	/*
	   db *sql.DB // database connection can be injected here.
	*/
}

type UserService struct {
	/*
		db *sql.DB // database connection can be injected here.
	*/
}

func (s *AuthService) AddUser(username, password string) error {
	if len(username) == 0 || len(password) == 0 {
		return errors.New("username or password cannot be empty")
	}
	// Some database related code to store user to database.
	return nil
}

func (s *UserService) EditProfile(username, email, name, picture string) error {
	if len(username) == 0 || len(email) == 0 || len(name) == 0 || len(picture) == 0 {
		return errors.New("invalid profile data")
	}
	// Some database related code to edit user profile.
	return nil
}
