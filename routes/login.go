package routes

import (
	"encoding/json"
	"net/http"

	"github.com/abdullahgumi/authentication/dao"

	U "github.com/abdullahgumi/authentication/models"
	"github.com/abdullahgumi/authentication/util"
	"github.com/julienschmidt/httprouter"
)

func Login(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

	user := U.User{}
	json.NewDecoder(req.Body).Decode(&user)
	loggedUser := dao.DoLogin(user.Username, user.Password)

	err1 := json.NewEncoder(res).Encode(loggedUser)
	util.HandleErr(err1)

}
