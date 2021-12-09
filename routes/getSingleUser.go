package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/abdullahgumi/authentication/dao"

	"github.com/abdullahgumi/authentication/util"
	"github.com/julienschmidt/httprouter"
)

func GetSingleUser(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params[0].Value)
	util.HandleErr(err)
	user := dao.DoFetchOneUser(id)
	err1 := json.NewEncoder(res).Encode(user)
	util.HandleErr(err1)
}
