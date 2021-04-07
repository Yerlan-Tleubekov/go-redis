package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yerlan-tleubekov/go-redis/pkg/response"

	"github.com/yerlan-tleubekov/go-redis/internal/models"
)

func (handler *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("it works")
}

func (handler *Handler) SignIn(w http.ResponseWriter, r *http.Request) {

	var userValid models.UserValidation = models.UserValidation{}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body, &userValid)

	if err != nil {
		fmt.Println(err)
		return
	}

	user, code, err := handler.services.GetUser(userValid.Login)

	if err != nil {
		response.WriteResponse(w, code, err)
		return
	}

	token, err, code := handler.services.SignIn(user.UserID)
	if err != nil {
		response.ErrorJsonWriter(err, code, w)
		return
	}

	userToken := response.Token{UserToken: token}
	answerJSON := response.AnswerJSON{Data: userToken, Code: code}
	response.WriteResponse(w, code, answerJSON)

}
