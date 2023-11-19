package models

import (
	"database/sql"
)

type Client struct {
	Login          string
	HashedPassword string
}

func GetClientByLogin(login string) (Client, error) {
	stmt, err := DB.Prepare("SELECT login, password from clients WHERE login = ?")

	if err != nil {
		return Client{}, err
	}

	client := Client{}

	sqlErr := stmt.QueryRow(login).Scan(&client.Login, &client.HashedPassword)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Client{}, nil
		}
		return Client{}, sqlErr
	}
	return client, nil
}
