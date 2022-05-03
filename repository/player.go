package repository

import (
	"cardGameSql/model"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

// CreatePlayer ...
func (s *sqlRepo) CreatePlayer(data *model.PlayerRegister) error {
	db := s.db

	sqlStmt := `INSERT INTO "players" (name, email, password, createdAt, updatedAt)
					 VALUES ($1, $2, $3, $4, $5)`

	err := db.QueryRow(sqlStmt, data.Name, data.Email, data.Password, time.Now(), time.Now())
	if err != nil {
		return err.Err()
	}

	return nil
}

func (s *sqlRepo) GetAllPlayer() ([]model.Player, error) {
	db := s.db
	players := make([]model.Player, 0)

	rows, err := db.Query(`SELECT * FROM "players"`)
	if err != nil {
		return players, err
	}

	defer rows.Close()
	for rows.Next() {
		player := model.Player{}
		// SELECT id, name, gender, email FROM public.users;
		err = rows.Scan(&player.ID, &player.Name, &player.Email, &player.Password, &player.CreatedAt, &player.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			break
		}

		players = append(players, player)
	}

	err = rows.Err()
	if err != nil {
		return players, err
	}

	return players, nil
}

func (s *sqlRepo) GetPlayerByID(id int) (model.Player, error) {
	db := s.db

	var player model.Player

	// create the select sql query
	sqlStmt := `SELECT * FROM players WHERE id=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStmt, id)

	// unmarshal the row object to user
	err := row.Scan(&player.ID, &player.Name, &player.Email, &player.Password, &player.CreatedAt, &player.UpdatedAt)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return player, nil
	case nil:
		return player, nil
	default:
		fmt.Println("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return player, err
}

func (s *sqlRepo) FindPlayerByEmail(email string) (*model.Player, error) {
	var player model.Player

	// sql query statement
	sqlStmt := `SELECT * from "players" WHERE email=$1`

	// execute the sql statement
	row := s.db.QueryRow(sqlStmt, email)

	// unmarshal the row object to user
	if err := row.Scan(&player.ID, &player.Name, &player.Email, &player.Password, &player.CreatedAt, &player.UpdatedAt); err != nil {
		return &player, err
	}

	return &player, nil
}
