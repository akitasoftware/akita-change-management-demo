package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type User struct {
	ID string `json:"id"`

	// Exclude phone from JSON responses... but what if we use YAML?
	Phone string `json:"-"`
}

func main() {
	myUsers := []User{
		{
			ID:    "usr_295oDMFK8b1yS5dwlSTdgP",
			Phone: "(123) 456-7890",
		},
		{
			ID:    "usr_6NiejyYEVpWfziUXJgovV6",
			Phone: "(777) 888-9999",
		},
	}

	http.HandleFunc("/users/json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		enc := json.NewEncoder(w)
		enc.Encode(myUsers)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
