package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/trevrosen/go-testing-presentation/webapp/db"
)

func App(cdb db.CommonDBInterface) http.Handler {
	routes := router(cdb)
	n := negroni.Classic()
	n.UseHandler(routes)
	return n
}

// router returns the router which defines the routes
func router(cdb db.CommonDBInterface) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users/{userID}", ShowUserHandler(cdb))
	return router
}

// ShowUserHandler returns the JSON representation of the User or a 404 header
func ShowUserHandler(cdb db.CommonDBInterface) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := mux.Vars(r)["userID"]
		user, err := cdb.FindUserById(userID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
			jsonData, _ := json.Marshal(&user)
			w.Write(jsonData)
		}
	})
}
