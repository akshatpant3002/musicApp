package main

import (
	"Golang-API/entity"
	groupRepository "Golang-API/groupRepository"
	userRepository "Golang-API/userRepository"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	groupRepo groupRepository.GroupRepository = groupRepository.NewGroupRepository()
	userRepo  userRepository.UserRepository   = userRepository.NewUserRepository()
)

/*
The getGroups function in route.go takes a response writer and request,
and then calls the FindAll method in groupRepository\groupPost-repo.go
to get all the groups in the group collection in the firestore database
*/
func getGroups(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	groups, err := groupRepo.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error gettings the groups"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(groups)
}

func handlePreflight(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	resp.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	resp.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	resp.WriteHeader(http.StatusOK)
}

/*
The getGroup function in route.go takes a response writer and request,
and then calls the FindGroup method in groupRepository\groupPost-repo.go
to get a specific group in the group collection in the firestore database
*/
func getGroup(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	vars := mux.Vars(req)
	groupID := vars["groupID"]
	log.Printf("Fetching group with ID: %s", groupID)
	group, err := groupRepo.FindGroup(groupID)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error getting the group"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(group)
}

/*
getUser function gets a specific user using the userID.
It uses the FindUser method in the UserRepository interface to fetch the user from the Firestore database.
If there is an error, it returns a 500 status code and an error message.
*/
func getUser(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	vars := mux.Vars(req)
	userID := vars["userID"]
	log.Printf("Fetching user with ID: %s", userID)
	user, err := userRepo.FindUser(userID)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error getting the user"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(user)
}

/*
The addGroups function in route.go takes a response writer and request,
and then calls the Save method in groupRepository\groupPost-repo.go
to add a group to the group collection in the firestore database
*/
func addGroups(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "OPTIONS" {
		handlePreflight(resp, req)
		return
	}

	resp.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	resp.Header().Set("Content-type", "application/json")
	var group entity.Group
	err := json.NewDecoder(req.Body).Decode(&group)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling the groups array"}`))
		return
	}
	groupRepo.Save(&group)
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(group)
}

/*
The getUsers function gets all the users from the userRepo and returns them as a JSON array.
It uses the FindAll method in the UserRepository interface to fetch the users from the Firestore database.
If there is an error, it returns a 500 status code and an error message.
*/
func getUsers(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	users, err := userRepo.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error gettings the users"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(users)
}

/*
addUsers function adds a new user to the userRepo.
It uses the Save method in the UserRepository interface to save the user to the Firestore database.
If there is an error, it returns a 500 status code and an error message.
*/
func addUsers(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "OPTIONS" {
		handlePreflight(resp, req)
		return
	}

	resp.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	resp.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	resp.Header().Set("Content-type", "application/json")

	if req.Method == "POST" {
		var user entity.User
		err := json.NewDecoder(req.Body).Decode(&user)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte(`{"error": "Error unmarshalling the users array"}`))
			return
		}
		userRepo.Save(&user)
		resp.WriteHeader(http.StatusOK)
		json.NewEncoder(resp).Encode(user)
	} else {
		resp.WriteHeader(http.StatusMethodNotAllowed)
		resp.Write([]byte(`{"error": "Method not allowed"}`))
	}
}
