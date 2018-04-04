package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/citrus-tart/certificate-aggregator/repository"
	"github.com/gorilla/mux"
)

var repo repository.Repository

func Init(r repository.Repository) {
	repo = r

	router := mux.NewRouter()
	router.HandleFunc("/certificates/{id}", certificatesHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func certificatesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result := repo.Certs[params["id"]]

	if result.ID != "" {
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 Certificate not found")
	}
}
