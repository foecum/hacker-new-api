package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/foecum/going/server"
	"github.com/gorilla/mux"
)

// RegisterRoutes regsiters routes for items
func (u User) RegisterRoutes(srv *server.Server) {
	router := srv.SetPathPrefix("/user")
	router.HandleFunc("/{username:[A-Za-z0-9]+}", getUserByUsername)
}

func getUserByUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var username string
	if v, ok := vars["username"]; ok {
		username = v
	}
	hnUser, err := getHackerNewUser(username)
	if err != nil {
		log.Printf("error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	b, err := json.Marshal(hnUser)
	if err != nil {
		log.Printf("error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(b))
}
