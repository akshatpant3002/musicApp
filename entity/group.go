package entity

type Group struct {
	GroupID string   `json:"groupID"`
	Users   []string `json:"users"`
}
