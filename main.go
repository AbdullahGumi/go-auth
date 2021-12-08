package main

import (
	"fmt"
	"net/http"

	"github.com/abdullahgumi/authentication/dao"

	"github.com/abdullahgumi/authentication/routes"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	dao.Init()

	router := httprouter.New()

	router.GET("/user/", routes.GetAllUsers)
	router.GET("/user/:id", routes.GetSingleUser)
	router.POST("/register/", routes.Register)
	router.POST("/login/", routes.Login)
	fmt.Println("Server listening to port 3001")

	handler := cors.AllowAll().Handler(router)
	http.ListenAndServe("localhost:3001", handler)
}
