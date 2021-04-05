package models

type User struct {
	UserID   int    `json:"userID"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserToken struct {
	UserID int    `json:"userID"`
	Token  string `json:"token"`
}
