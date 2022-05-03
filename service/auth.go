package service

import (
	"cardGameSql/model"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// AuthRepo ...
type AuthRepo interface {
	FindPlayerByEmail(email string) (*model.Player, error)
	CreatePlayer(data *model.PlayerRegister) error
}

// authService ...
type authService struct {
	repo AuthRepo
}

// NewAuthService ...
func NewAuthService(repo AuthRepo) *authService {
	return &authService{repo: repo}
}

// Register ...
func (as *authService) Register(data *model.PlayerRegister) error {

	// FindPlayerByEmail
	player, err := as.repo.FindPlayerByEmail(data.Email)
	if err != nil {
		return err
	}

	// email has already existed
	if player != nil {
		return errors.New("email has already existed")
	}

	// hash player password
	bytes, err := bcrypt.GenerateFromPassword([]byte(player.Password), 14)
	if err != nil {
		return err
	}
	data.Password = string(bytes)

	// call repo
	if err = as.repo.CreatePlayer(data); err != nil {

		return err
	}

	return nil
}
