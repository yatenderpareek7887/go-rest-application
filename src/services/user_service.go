package services

import (
	"errors"
	"go-rest-application/src/config/storage"
	userdto "go-rest-application/src/dto/user"
	models "go-rest-application/src/models/user"
	"os"
	"strconv"
)

func GetAllUsers() ([]models.User, error) {
	return storage.LoadUsers()
}

func CreateUser(userDTO userdto.CreateUserRequest) (models.User, error) {
	users, err := storage.LoadUsers()
	if err != nil && !os.IsNotExist(err) {
		return models.User{}, err
	}

	user := models.User{
		ID:    len(users) + 1,
		Name:  userDTO.Name,
		Email: userDTO.Email,
		City:  userDTO.City,
		Age:   userDTO.Age,
	}

	users = append(users, user)
	if err := storage.SaveUsers(users); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func UpdateUser(id string, updatedUser userdto.CreateUserRequest) (models.User, error) {
	users, err := storage.LoadUsers()
	if err != nil {
		return models.User{}, err
	}
	for i, user := range users {
		if strconv.Itoa(user.ID) == id {
			users[i].Name = updatedUser.Name
			users[i].Email = updatedUser.Email
			users[i].City = updatedUser.City
			users[i].Age = updatedUser.Age

			if err := storage.SaveUsers(users); err != nil {
				return models.User{}, err
			}

			return users[i], nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func UpdatePartialUser(id string, partialUser userdto.UpdateUserRequest) (models.User, error) {
	users, err := storage.LoadUsers()
	if err != nil {
		return models.User{}, err
	}
	for i, user := range users {
		if strconv.Itoa(user.ID) == id {
			if partialUser.Name != "" {
				users[i].Name = partialUser.Name
			}
			if partialUser.Email != "" {
				users[i].Email = partialUser.Email
			}
			if partialUser.Age != 0 {
				users[i].Age = partialUser.Age
			}
			if partialUser.City != "" {
				users[i].City = partialUser.City
			}

			if err := storage.SaveUsers(users); err != nil {
				return models.User{}, err
			}
			return users[i], nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func DeleteUser(id string) error {
	users, err := storage.LoadUsers()
	if err != nil {
		return err
	}
	for i, user := range users {
		if strconv.Itoa(user.ID) == id {
			users = append(users[:i], users[i+1:]...)
			if err := storage.SaveUsers(users); err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("user not found")
}
