package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/foecum/hackernews/item"
	"github.com/foecum/hackernews/server"
)

func main() {
	srv := server.NewServer("", "8080")

	srv.RegisterHandler("/items", func(w http.ResponseWriter, r *http.Request) {
		queries, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			log.Printf("error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		itemIDStr := queries.Get("id")
		if itemIDStr == "" {
			log.Printf("error: %v", err)
			w.WriteHeader(http.StatusNotFound)
			return
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
		item, err := item.GetHackerNewItem(itemID)
		if err != nil {
			log.Printf("error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		b, err := json.Marshal(item)
		if err != nil {
			log.Printf("error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, string(b))
	})

	srv.StartServer()
}
