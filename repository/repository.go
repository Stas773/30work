package repository

import (
	"D/30work/models"
	"fmt"

	"github.com/google/uuid"
)

var Repository = make(map[string]models.User, 0)

type RepoStruct struct {
}

func (rs *RepoStruct) CreateNewUser(mu models.User) error {
	stringId := uuid.NewString()
	mu.Id = stringId

	Repository[mu.Id] = mu
	return nil
}

func (rs *RepoStruct) MakeFriends(id1 string, id2 string) (error, error) {
	user1, ok := Repository[id1]
	if !ok {
		err := fmt.Errorf("user1 not found")
		return err, nil
	}

	user2, ok := Repository[id2]
	if !ok {
		err := fmt.Errorf("user2 not found")
		return nil, err
	}

	user1.Friends = append(user1.Friends, user2.Id)
	user2.Friends = append(user2.Friends, user1.Id)

	Repository[user1.Id] = user1
	Repository[user2.Id] = user2
	return nil, nil
}

func (rs *RepoStruct) GetUser(id string) (models.User, error) {
	user, ok := Repository[id]
	if !ok {
		err := fmt.Errorf("user not found")
		return user, err
	}
	return user, nil
}

func (rs *RepoStruct) GetAllUsers() (map[string]models.User, error) {
	allUsers := Repository
	if len(allUsers) == 0 {
		err := fmt.Errorf("users not found")
		return allUsers, err
	}
	return allUsers, nil
}

func (rs *RepoStruct) GetUserFriends(id string) ([]string, error) {
	user, ok := Repository[id]
	if !ok {
		err := fmt.Errorf("users not found")
		return user.Friends, err
	}
	return user.Friends, nil
}

func (rs *RepoStruct) UpdateUser(id string, mu models.User) error {
	user, ok := Repository[id]
	if !ok {
		err := fmt.Errorf("users not found")
		return err
	}

	if mu.Age > 0 {
		user.Age = mu.Age
	}
	if mu.Name != "" {
		user.Name = mu.Name
	}
	Repository[user.Id] = user
	return nil
}
func (rs *RepoStruct) DeleteUser(id string) error {
	user, ok := Repository[id]
	if !ok {
		err := fmt.Errorf("user not found")
		return err
	}

	for _, friendId := range user.Friends {
		friend, ok := Repository[friendId]
		if !ok {
			continue
		}
		for i, friendId := range friend.Friends {
			if friendId == id {
				friend.Friends = append(friend.Friends[:i], friend.Friends[i+1:]...)
			}
		}
		Repository[friend.Id] = friend
	}
	delete(Repository, id)
	return nil
}
