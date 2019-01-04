package controllers

import (
	"net/http"

	mgo "gopkg.in/mgo.v2"
)

// ManagerController ...
type ManagerController struct {
	session *mgo.Session
}

// NewManagerController ...
func NewManagerController(session *mgo.Session) *ManagerController {
	return &ManagerController{
		session: session,
	}
}

// Register ...
func (mc *ManagerController) Register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		Allo tu pues	
	`))
}
