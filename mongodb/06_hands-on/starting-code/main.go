package main

import (
	"github.com/RemLampa/go-webdev/mongodb/06_hands-on/starting-code/controllers"
	"github.com/RemLampa/go-webdev/mongodb/06_hands-on/starting-code/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var users map[string]models.User

func main() {
	users = make(map[string]models.User)

	r := httprouter.New()

	// Get a UserController instance
	uc := controllers.NewUserController(users)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}
