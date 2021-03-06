package main

import (
	"GoBasic/StrukturaProjekta/API"
	"github.com/gin-gonic/gin"
)

//Naredimo Router objekt da na njega obesimo funkcije
type Router struct {
	engine *gin.Engine
	api    API.Controller
}

func (r *Router) registerRoutes() (err error) {

	//Pot /api/v1
	api := r.engine.Group("/api/v1")

	//Pot /api/v1/user
	user := api.Group("/user")
	r.registerUserRoutes(user)

	return

}

func (r *Router) registerUserRoutes(user *gin.RouterGroup) {

	//Pot /api/v1/user
	user.GET("/", r.api.GetUsers)

	//Pot /api/v1/user/:user_id
	user.GET("/:user_id", r.api.GetUserById)

}
