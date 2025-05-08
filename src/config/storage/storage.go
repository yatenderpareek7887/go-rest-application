package storage

import (
	"encoding/json"
	models "go-rest-application/src/models/user"
	"os"
)

var userFile = "../DB/users.json"

func SaveUsers(users []models.User) error {
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(userFile, data, 0644)
}

func LoadUsers() ([]models.User, error) {
	file, err := os.ReadFile(userFile)
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err := json.Unmarshal(file, &users); err != nil {
		return nil, err
	}
	return users, nil
}
