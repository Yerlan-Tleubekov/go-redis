package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yerlan-tleubekov/go-redis/pkg/response"

	"github.com/yerlan-tleubekov/go-redis/internal/models"
	// "github.com/yerlan-tleubekov/go-redis/internal/cache"
)

// var (
// tokenCache cache.TokenCache
// )

func (handler *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("it works")
}

func (handler *Handler) SignIn(w http.ResponseWriter, r *http.Request) {

	var userIdent models.UserIndent = models.UserIndent{}

	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body, &userIdent)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(userIdent.UserID)

	token, err, code := handler.services.SignIn(userIdent.UserID)
	if err != nil {
		response.ErrorJsonWriter(err, code, w)
		return
	}

	userToken := response.Token{UserToken: token}
	answerJSON := response.AnswerJSON{Data: userToken}
	response.WriteResponse(w, code, answerJSON)

}
