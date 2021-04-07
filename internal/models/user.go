package models

type UserIndent struct {
	UserID int `json:"userID"`
}

type User struct {
	UserID   int    `json:"userID"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserValidation struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserToken struct {
	UserID int    `json:"userID"`
	Token  string `json:"token"`
}
