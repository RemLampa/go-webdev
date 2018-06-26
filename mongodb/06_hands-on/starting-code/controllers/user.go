package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/RemLampa/go-webdev/mongodb/06_hands-on/starting-code/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

// UserController is a struct
type UserController struct {
	users map[bson.ObjectId]models.User
}

// NewUserController returns a new user controller
func NewUserController(users map[bson.ObjectId]models.User) *UserController {
	return &UserController{users}
}

// GetUser fetches user from database and writes it on the response
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")
	// Verify id is ObjectId hex representation, otherwise return status not found
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	// ObjectIdHex returns an ObjectId from the provided hex representation.
	// oid := bson.ObjectIdHex(id)

	// composite literal
	u := models.User{}

	// Fetch user
	//	if err := uc.session.DB("go-web-dev-db").C("users").FindId(oid).One(&u); err != nil {
	//		w.WriteHeader(404)
	//		return
	//	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

// CreateUser adds a new user to the database
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	// create bson ID
	u.Id = bson.NewObjectId()

	uc.users[u.Id] = u

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

// DeleteUser removes a user from the database
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	// Delete user
	//if err := uc.session.DB("go-web-dev-db").C("users").RemoveId(oid); err != nil {
	//	w.WriteHeader(404)
	//	return
	//}

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", oid, "\n")
}
