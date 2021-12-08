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

func Login(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

	user := U.User{}
	json.NewDecoder(req.Body).Decode(&user)
	loggedUser := dao.DoLogin(user.Username, user.Password)

	err1 := json.NewEncoder(res).Encode(U.Response{SignedTokenString(loggedUser), loggedUser})
	util.HandleErr(err1)

}

func SignedTokenString(user U.User) string {
	var st string

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"password": user.Password,
	})

	tokenString, err := token.SignedString([]byte(st))
	if err != nil {
		log.Println(err)
	}
	return tokenString
}
