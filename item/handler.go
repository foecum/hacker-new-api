package item

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/foecum/going/server"
	"github.com/gorilla/mux"
)

// RegisterRoutes regsiters routes for items
func (item Item) RegisterRoutes(srv *server.Server) {
	router := srv.SetPathPrefix("/item")
	router.HandleFunc("/{id:[0-9]+}", getItemByID)
}

func getItemByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var itemIDStr string
	if v, ok := vars["id"]; ok {
		itemIDStr = v
	}

	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		log.Printf("error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if itemID == 0 {
		log.Printf("error: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	hnItem, err := getHackerNewItem(itemID)
	if err != nil {
		log.Printf("error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	b, err := json.Marshal(hnItem)
	if err != nil {
		log.Printf("error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(b))
}
