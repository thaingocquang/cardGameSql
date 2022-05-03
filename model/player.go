package model

import (
	"errors"
	"time"
)

type (
	// Player ...
	Player struct {
		ID        int
		Name      string
		Email     string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	// PlayerRegister ...
	PlayerRegister struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

// Validate ...
func (p PlayerRegister) Validate() error {
	if len(p.Name) == 0 {
		return errors.New("name can't be empty")
	}
	if len(p.Email) == 0 {
		return errors.New("email can't be empty")
	}
	if len(p.Password) == 0 {
		return errors.New("password can't be empty")
	}
	return nil
}
