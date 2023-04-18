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
	resp.Header().Set("Content-type", "application/json")
	var user entity.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling the users array"}`))
		return
	}
	//gets all the current users to verify that any user being added to friends list is actually a user
	users, err := userRepo.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error gettings the users"}`))
		return
	}
	friendsExist := true
	// Check here to see if the friends are actually users.
	for _, friendID := range user.Friends {
		friendExists := false
		for _, u := range users {
			if u.UserID == friendID {
				friendExists = true
				break
			}
		}
		if !friendExists {
			friendsExist = false
			break
		}
	}
	if !friendsExist {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(`{"error": "One or more friends do not exist"}`))
		return
	}
	userRepo.Save(&user)
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(user)
}

/*
putUsers function updates a pre-existing user to the userRepo.
It uses the Save method in the UserRepository interface to save the user to the Firestore database.
If there is an error, it returns a 500 status code and an error message.
*/
func putUsers(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var user entity.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling the users array"}`))
		return
	}
	//gets all the current users to verify that any user being added to friends list is actually a user
	users, err := userRepo.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error gettings the users"}`))
		return
	}
	friendsExist := true
	// Check here to see if the friends are actually users.
	for _, friendID := range user.Friends {
		friendExists := false
		for _, u := range users {
			if u.UserID == friendID {
				friendExists = true
				break
			}
		}
		if !friendExists {
			friendsExist = false
			break
		}
	}
	if !friendsExist {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(`{"error": "One or more friends do not exist"}`))
		return
	}
	userRepo.Update(&user)
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(user)
}

/*
the PutGroups function updates a pre-existing group to the groupRepo.
It uses the Update method in the GroupRepository interface to save the group to the Firestore database.
If there is an error, it returns a 500 status code and an error message.
*/

func putGroup(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var group entity.Group // The JSON body is decoded into this variable.
	err := json.NewDecoder(req.Body).Decode(&group)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling the groups array"}`))
		return
	}

	//gets all the current users to verify that any user being added to friends list is actually a user
	users, err := userRepo.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error gettings the users"}`))
		return
	}

	// Check here to see if the members are actually users.
	membersExist := true
	for _, memberID := range group.Users {
		memberExists := false
		for _, u := range users {
			if u.UserID == memberID {
				memberExists = true
				break
			}
		}
		if !memberExists {
			membersExist = false
			break
		}
	}
	if !membersExist {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(`{"error": "One or more members do not exist"}`))
		return
	}
	groupRepo.Update(&group)
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(group)
}

func updatePlaylistName(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var group entity.Group
	err := json.NewDecoder(req.Body).Decode(&group)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling the groups array"}`))
		return
	}
	groupFromRepo, err := groupRepo.FindGroup(group.GroupID)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error getting the group"}`))
		return
	}

	group.Matched = groupFromRepo.Matched
	group.SemiMatched = groupFromRepo.SemiMatched
	group.Users = groupFromRepo.Users

	groupRepo.Update(&group)
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(groupFromRepo)
}

func updateGroupUsers(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var group entity.Group
	err := json.NewDecoder(req.Body).Decode(&group)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling the groups array"}`))
		return
	}

	users, err := userRepo.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error gettings the users"}`))
		return
	}

	// Check here to see if the members are actually users.
	membersExist := true
	for _, memberID := range group.Users {
		memberExists := false
		for _, u := range users {
			if u.UserID == memberID {
				memberExists = true
				break
			}
		}
		if !memberExists {
			membersExist = false
			break
		}
	}
	if !membersExist {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(`{"error": "One or more members do not exist"}`))
		return
	}

	//gets the group from the database
	groupFromDB, err := groupRepo.FindGroup(group.GroupID)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error getting the group"}`))
		return
	}

	//adds the new users to the group
	for _, user := range groupFromDB.Users {
		//log.Printf("User: %s", user)
		group.Users = append(group.Users, user)
	}

	//log.Printf("Group: %s", group)

	//removes any duplicates
	group.Users = removeDuplicates(group.Users)

	groupRepo.Update(&group)
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(group)
}

func removeDuplicates(s []string) []string {
	seen := make(map[string]bool)
	j := 0
	for i, v := range s {
		if seen[v] {
			continue
		}
		seen[v] = true
		s[j] = s[i]
		j++
	}
	s = s[:j]

	return s
}

func deleteUser(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	vars := mux.Vars(req)
	userID := vars["userID"]
	log.Printf("Deleting user with ID: %s", userID)
	err := userRepo.DeleteUser(userID)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error deleting the user"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode("User deleted successfully")
}

func deleteGroup(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	vars := mux.Vars(req)
	groupID := vars["groupID"]
	log.Printf("Deleting group with ID: %s", groupID)
	err := groupRepo.DeleteGroup(groupID)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error deleting the group"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode("Group deleted successfully")
}

func updateUserGroups(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var user entity.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling the groups array"}`))
		return
	}

	userFromDB, err := userRepo.FindUser(user.UserID)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error getting the user"}`))
		return
	}

	//adds the new groups to the user

	for _, group := range userFromDB.Groups {
		user.Groups = append(user.Groups, group)
	}

	user.Groups = removeDuplicates(user.Groups)

	//user.Groups = removeDuplicates(user.Groups)

	// keep old user data for fields not being updated
	user.Friends = userFromDB.Friends
	user.LikedSong = userFromDB.LikedSong
	user.GroupAdmin = userFromDB.GroupAdmin

	userRepo.Update(&user)
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(user)
}

func updateUserFriends(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var user entity.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling the friends array"}`))
		return
	}

	users, err := userRepo.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error getting the users"}`))
		return
	}

	// Check here to see if the members are actually users.
	membersExist := true
	for _, memberID := range user.Friends {
		memberExists := false
		for _, u := range users {
			if u.UserID == memberID {
				memberExists = true
				break
			}
		}
		if !memberExists {
			membersExist = false
			break
		}
	}
	if !membersExist {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(`{"error": "One or more members do not exist"}`))
		return
	}

	userFromDB, err := userRepo.FindUser(user.UserID)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error getting the user"}`))
		return
	}

	//adds the new friends to the user
	for _, friend := range userFromDB.Friends {
		user.Friends = append(user.Friends, friend)
	}

	//removes any duplicates
	user.Friends = removeDuplicates(user.Friends)

	// keep old user data for fields not being updated
	user.LikedSong = userFromDB.LikedSong
	user.GroupAdmin = userFromDB.GroupAdmin
	user.Groups = userFromDB.Groups

	userRepo.Update(&user)
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(user)
}

func addGroupLikedSong(resp http.ResponseWriter, req *http.Request) {

	resp.Header().Set("Content-type", "application/json")
	vars := mux.Vars(req)
	groupID := vars["groupID"]

	var song string
	err := json.NewDecoder(req.Body).Decode(&song)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling the song"}`))
		return
	}

	group, err := groupRepo.FindGroup(groupID)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error getting the group"}`))
		return
	}

	//adds the new song to the group
	group.Matched = append(group.Matched, song)

	//removes any duplicates
	group.Matched = removeDuplicates(group.Matched)

	groupRepo.Update(group)
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(group)

}

func checkIfMatched(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	vars := mux.Vars(req)
	groupID := vars["groupID"]
	songID := vars["songID"]
	groupFromDB, err := groupRepo.FindGroup(groupID)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error getting the group"}`))
		return
	}
	isInMatched := false

	//checks if the group.Matched has the songID
	for _, song := range groupFromDB.Matched {
		if song == songID {
			isInMatched = true
			break
		}
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(isInMatched)
}
