package main

import (
	"Golang-API/entity"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetGroups(t *testing.T) {
	req, err := http.NewRequest("GET", "/groupPost", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getGroups)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expectedGroups := make([]entity.Group, 0)
	if err := json.Unmarshal(rr.Body.Bytes(), &expectedGroups); err != nil {
		t.Fatal(err)
	}
	if len(expectedGroups) == 0 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), "[]")
	}
}

func TestAddGroups(t *testing.T) {
	group := entity.Group{
		GroupID:     "finalSprintGroupTesting",
		Users:       []string{"akshatpant3002"},
		PlaylistID:  "test-playlist",
		SemiMatched: map[string]float64{"user1": 0.5, "user2": 0.7},
		Matched:     []string{"song1", "song2"},
	}

	body, err := json.Marshal(group)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/groupPost", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(addGroups)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := entity.Group{}
	json.Unmarshal(rr.Body.Bytes(), &expected)
	if expected.GroupID != "finalSprintGroupTesting" {
		t.Errorf("handler returned unexpected body: got %v want group with ID test-groupforsprint3",
			expected)
	}
}

func TestGetGroup(t *testing.T) {
	req, err := http.NewRequest("GET", "/groupPost/{groupID}", nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{"groupID": "finalSprintGroupTesting"}
	req = mux.SetURLVars(req, vars)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getGroup)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expectedGroup := entity.Group{}
	if err := json.Unmarshal(rr.Body.Bytes(), &expectedGroup); err != nil {
		t.Fatal(err)
	}
	if expectedGroup.GroupID != "finalSprintGroupTesting" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), `{"groupID":"finalSprintGroupTesting","users":["user1","user2"],"playlistID":"test-playlist","semiMatched":{"user1":0.5,"user2":0.7},"matched":{"user1":131.4,"user2":0.9}}`)
	}
}

func TestPutGroup(t *testing.T) {
	group := entity.Group{
		GroupID:     "finalSprintGroupTesting",
		Users:       []string{"akshatpant3002"},
		PlaylistID:  "test-playlist",
		SemiMatched: map[string]float64{"songPUT1": 1, "songPUT2": 1},
		Matched:     []string{"song1", "song2"},
	}

	body, err := json.Marshal(group)
	if err != nil {
		t.Fatal(err)
	}

	req1, err1 := http.NewRequest("PUT", "/groupPost/{groupID}", strings.NewReader(string(body)))
	if err1 != nil {
		t.Fatal(err1)
	}
	vars := map[string]string{"groupID": "finalSprintGroupTesting"}
	req1 = mux.SetURLVars(req1, vars)
	rr1 := httptest.NewRecorder()
	handler := http.HandlerFunc(putGroup)
	handler.ServeHTTP(rr1, req1)
	if status := rr1.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := entity.Group{}
	json.Unmarshal(rr1.Body.Bytes(), &expected)
	if expected.GroupID != "finalSprintGroupTesting" {
		t.Errorf("handler returned unexpected body: got %v want group with ID test-groupforsprint3",
			expected)
	}
}

func TestUpdateGroupUsers(t *testing.T) {
	group := entity.Group{
		GroupID: "finalSprintGroupTesting",
		Users:   []string{"akshatpant3002", "31furtgoascpojlk76vcalfbyqle"},
	}

	body, err := json.Marshal(group)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/groupPost/{groupID}", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{"groupID": "finalSprintGroupTesting"}
	req = mux.SetURLVars(req, vars)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updateGroupUsers)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := entity.Group{}
	json.Unmarshal(rr.Body.Bytes(), &expected)
	if expected.Users[1] != "31furtgoascpojlk76vcalfbyqle" {
		t.Errorf("handler returned unexpected body: got %v want group with ID finalSprintGroupTesting",
			expected)
	}
}

func TestUpdatePlaylistName(t *testing.T) {
	group := entity.Group{
		GroupID:    "finalSprintGroupTesting",
		PlaylistID: "updated-name",
	}

	body, err := json.Marshal(group)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/userUpdatePlaylists/", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updatePlaylistName)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := entity.Group{}
	json.Unmarshal(rr.Body.Bytes(), &expected)

	if expected.GroupID != "finalSprintGroupTesting" {
		t.Errorf("handler returned unexpected body: got %v want group with playlistID update-name",
			expected)
	}
}

func TestAddGroupLikedSong(t *testing.T) {
	songID := "testAddID"

	body, err := json.Marshal(songID)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/groupPost/{groupID}", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{"groupID": "finalSprintGroupTesting"}
	req = mux.SetURLVars(req, vars)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(addGroupLikedSong)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := entity.Group{}
	json.Unmarshal(rr.Body.Bytes(), &expected)
	if expected.Matched[2] != songID {
		t.Errorf("handler returned unexpected body: got %v want testAddID with value 1",
			expected)
	}
}

func TestCheckIfMatched(t *testing.T) {
	req, err := http.NewRequest("GET", "/checkIfMatched/{groupID}/{songID}", nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := map[string]string{"groupID": "finalSprintGroupTesting", "songID": "testAddID"}
	req = mux.SetURLVars(req, vars)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(checkIfMatched)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	matched := false
	if err := json.Unmarshal(rr.Body.Bytes(), &matched); err != nil {
		t.Fatal(err)
	}
	if !matched {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), matched)
	}
}

func TestDeleteGroup(t *testing.T) {
	//first get all the groups
	req, err := http.NewRequest("GET", "/groupPost", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getGroups)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expectedGroups := make([]entity.Group, 0)
	if err := json.Unmarshal(rr.Body.Bytes(), &expectedGroups); err != nil {
		t.Fatal(err)
	}

	//then delete one of them
	req1, err1 := http.NewRequest("DELETE", "/groupPost/{groupID}", nil)
	if err1 != nil {
		t.Fatal(err1)
	}
	vars := map[string]string{"groupID": "finalSprintGroupTesting"}
	req1 = mux.SetURLVars(req1, vars)
	rr1 := httptest.NewRecorder()
	handler = http.HandlerFunc(deleteGroup)
	handler.ServeHTTP(rr1, req1)
	if status := rr1.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	//then get all the groups again
	req, err = http.NewRequest("GET", "/groupPost", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(getGroups)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expectedGroupsAfterDelete := make([]entity.Group, 0)
	if err := json.Unmarshal(rr.Body.Bytes(), &expectedGroupsAfterDelete); err != nil {
		t.Fatal(err)
	}
	//then check the length is one less
	if len(expectedGroups)-len(expectedGroupsAfterDelete) != 1 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), "[]")
	}
}
