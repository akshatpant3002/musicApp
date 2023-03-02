package main

import (
	"Golang-API/entity"
	groupRepository "Golang-API/groupRepository"
	userRepository "Golang-API/userRepository"
	"encoding/json"
	"net/http"
)

var (
	groupRepo groupRepository.GroupRepository = groupRepository.NewGroupRepository()
	userRepo  userRepository.UserRepository   = userRepository.NewUserRepository()
)

//group get and add

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

// user get and add
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
