package response

type AnswerJSON struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
}

type Token struct {
	UserToken string `json:"token"`
}
