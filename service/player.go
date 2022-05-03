package service

import (
	"cardGameSql/model"
	"context"
)

// PlayerRepo ...
type PlayerRepo interface {
	CreatePlayer(data *model.PlayerRegister) error
	GetAllPlayer() ([]model.Player, error)
	GetPlayerByID(id int) (model.Player, error)
}

// playerService ...
type playerService struct {
	repo PlayerRepo
}

// NewPlayerService ...
func NewPlayerService(repo PlayerRepo) *playerService {
	return &playerService{repo: repo}
}

// CreateUser ...
func (p *playerService) CreateUser(ctx context.Context, data *model.PlayerRegister) error {

	err := p.repo.CreatePlayer(data)
	if err != nil {
		return err
	}

	return nil
}

// GetAllPlayer ...
func (p *playerService) GetAllPlayer() ([]model.Player, error) {

	players, err := p.repo.GetAllPlayer()
	if err != nil {
		return players, err
	}

	return players, nil
}

// GetPlayerByID ...
func (p *playerService) GetPlayerByID(id int) (model.Player, error) {

	player, err := p.repo.GetPlayerByID(id)
	if err != nil {
		return player, err
	}

	return player, nil
}
