package usecase

import (
	"D/30work/models"
)

type RepoInterface interface {
	CreateNewUser(user models.User) error
	GetUser(id string) (models.User, error)
	GetAllUsers() (map[string]models.User, error)
	MakeFriends(id1 string, id2 string) (error, error)
	UpdateUser(id string, mu models.User) error
	DeleteUser(id string) error
	GetUserFriends(id string) ([]string, error)
}

func CreateNewUser(r RepoInterface, mu models.User) error {
	return r.CreateNewUser(mu)
}

func MakeFriends(r RepoInterface, id1 string, id2 string) (error, error) {
	return r.MakeFriends(id1, id2)
}

func GetUser(r RepoInterface, id string) (models.User, error) {
	return r.GetUser(id)
}

func GetAllUsers(r RepoInterface) (map[string]models.User, error) {
	return r.GetAllUsers()
}

func GetUserFriends(r RepoInterface, id string) ([]string, error) {
	return r.GetUserFriends(id)
}

func UpdateUser(r RepoInterface, id string, mu models.User) error {
	return r.UpdateUser(id, mu)
}

func DeleteUser(r RepoInterface, id string) error {
	return r.DeleteUser(id)
}
