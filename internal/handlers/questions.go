package handlers

import (
	"back-sabervest/internal/models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

var DB *sql.DB
var questions models.Test

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	bks, err := questions.AllQuestions()
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	//for _, bk := range bks {
	//	fmt.Fprintf(w, "%v, %s, %v, ", bk.ID, bk.Name, bk.LocationId)
	//}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(bks)
}

func GetQuestionsParams(w http.ResponseWriter, r *http.Request) {
	topic := r.URL.Query().Get("topic")
	matter := r.URL.Query().Get("matter")
	dateInitial := r.URL.Query().Get("date_initial")
	dateFinal := r.URL.Query().Get("date_final")

	pageStr := r.URL.Query().Get("page")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		http.Error(w, "error is convert", http.StatusBadRequest)
		return
	}

	offset := page
	limit := offset - page

	if topic == "" || matter == "" || dateInitial == "" {
		http.Error(w, "Params is required", http.StatusBadRequest)
		return
	}

	filterQuestions := models.FilterQuestions{
		Topic:       topic,
		Matter:      matter,
		InitialDate: dateInitial,
		FinalDate:   dateFinal,
		Limit:       limit,
		Offset:      offset,
	}

	bks, err := questions.QuestionsParams(filterQuestions)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	//for _, bk := range bks {
	//	fmt.Fprintf(w, "%v, %s, %v, ", bk.ID, bk.Name, bk.LocationId)
	//}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(bks)
}
