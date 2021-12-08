package routes

import (
	"encoding/json"
	"net/http"

	"github.com/abdullahgumi/authentication/dao"
	U "github.com/abdullahgumi/authentication/models"
	"github.com/julienschmidt/httprouter"
)

type Response struct {
	Message string `json:"message"`
}

func Register(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

	newUser := U.User{}
	json.NewDecoder(req.Body).Decode(&newUser)
	isReg := dao.DoRegister(newUser.Username, newUser.Email, newUser.Password)

	if isReg != true {
		json.NewEncoder(res).Encode(Response{"An error occured"})
	} else {
		json.NewEncoder(res).Encode(Response{"User created "})
	}

}
