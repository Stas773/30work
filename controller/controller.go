package main

import (
	"D/30work/models"
	"D/30work/repository"
	"D/30work/usecase"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	router.Post("/users", CreateNewUser)
	router.Get("/users/{userId}", GetUser)
	router.Get("/users", GetAllUsers)
	router.Get("/users/{userId}/friends", GetUserFriends)
	router.Put("/users/{userId}/friends", MakeFriends)
	router.Patch("/users/{userId}", UpdateUser)
	router.Delete("/users/{userId}", DeleteUser)

	http.ListenAndServe(":8080", router)
}

func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var modelUser models.User

	err := json.NewDecoder(r.Body).Decode(&modelUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if modelUser.Name == "" {
		http.Error(w, "invalid name", http.StatusBadRequest)
		return
	}
	if modelUser.Age == 0 {
		http.Error(w, "age must be greater than 0", http.StatusBadRequest)
		return
	}

	err = usecase.CreateNewUser(&repository.RepoStruct{}, modelUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func MakeFriends(w http.ResponseWriter, r *http.Request) {
	type FriendRequest struct {
		User1 string `json:"userId1"`
		User2 string `json:"userId2"`
	}

	var friendRequest FriendRequest

	err := json.NewDecoder(r.Body).Decode(&friendRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err1, err2 := usecase.MakeFriends(&repository.RepoStruct{}, friendRequest.User1, friendRequest.User2)
	if err1 != nil {
		http.Error(w, "User1 not found", http.StatusNotFound)
		return
	} else if err2 != nil {
		http.Error(w, "User2 not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User %s and %s are friends now", friendRequest.User1, friendRequest.User2)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")

	user, err := usecase.GetUser(&repository.RepoStruct{}, userId)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := usecase.GetAllUsers(&repository.RepoStruct{})
	if err != nil {
		http.Error(w, "Users not found", http.StatusNotFound)
		return
	}
	for _, v := range users {
		json.NewEncoder(w).Encode(v)
	}
}

func GetUserFriends(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")

	userFriends, err := usecase.GetUserFriends(&repository.RepoStruct{}, userId)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(userFriends)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	var updateUser models.User
	err := json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = usecase.UpdateUser(&repository.RepoStruct{}, userId, updateUser)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User %s updated", userId)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	err := usecase.DeleteUser(&repository.RepoStruct{}, userId)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User %s deleted", userId)
}
