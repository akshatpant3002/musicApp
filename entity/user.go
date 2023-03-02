package entity

// structure for user containing friends, liked songs and user id
// friends and likedsong should contain an array of strings and userID should be a string
type User struct {
	Friends    []string        `json:"friends"`
	LikedSong  []string        `json:"likedSong"`
	UserID     string          `json:"userID"`
	GroupAdmin map[string]bool `json:"groupAdmin"` // key is group id, admin status
}
