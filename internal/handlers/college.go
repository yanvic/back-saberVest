package handlers

import (
	"back-sabervest/internal/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var university models.College

func GetUniversity(w http.ResponseWriter, r *http.Request) {
	bks, err := university.AllUniversity()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%v, %s, %v, ", bk.ID, bk.Name, bk.LocationId)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(bks)
}
