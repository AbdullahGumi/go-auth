package routes

import (
	"encoding/json"
	"net/http"

	"github.com/abdullahgumi/authentication/dao"

	"github.com/abdullahgumi/authentication/util"
	"github.com/julienschmidt/httprouter"
)

func GetAllUsers(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	users := dao.DoFetchAllUsers()
	err1 := json.NewEncoder(res).Encode(users)
	util.HandleErr(err1)
}
