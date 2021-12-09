package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/abdullahgumi/authentication/dao"

	U "github.com/abdullahgumi/authentication/models"
	"github.com/abdullahgumi/authentication/util"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

type LoginResponse struct {
	User    U.User `json:"user,omitempty"`
	Token   string
	Message string
	Status  bool
}

type BadLoginResponse struct {
	Message string
	Status  bool
}

func Login(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	user := U.User{}
	json.NewDecoder(req.Body).Decode(&user)
	loggedUser := dao.DoLogin(user.Email, user.Password)

	if loggedUser.FirstName != "" {
		token := SignedTokenString(loggedUser)
		message := "user logged in successfully"
		status := true
		response := LoginResponse{loggedUser, token, message, status}
		err1 := json.NewEncoder(res).Encode(response)
		util.HandleErr(err1)
	} else {
		message := "Invalid username or password"
		status := false
		response := BadLoginResponse{message, status}
		err1 := json.NewEncoder(res).Encode(response)
		util.HandleErr(err1)
	}

}

func SignedTokenString(user U.User) string {
	var st string
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
	})

	tokenString, err := token.SignedString([]byte(st))
	if err != nil {
		log.Println(err)
	}
	return tokenString
}
