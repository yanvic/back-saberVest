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
	topicId := r.URL.Query().Get("topic")
	matterId := r.URL.Query().Get("matter")
	universityId := r.URL.Query().Get("university")
	dateInitial := r.URL.Query().Get("date_initial")
	dateFinal := r.URL.Query().Get("date_final")

	pageStr := r.URL.Query().Get("page")

	//page, err := strconv.Atoi(pageStr)
	//if err != nil {
	//	http.Error(w, "error is convert", http.StatusBadRequest)
	//	return
	//}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		log.Printf("Parâmetro de página inválido: %v, usando 1 como padrão", pageStr)
		page = 1
	}

	topic, err := strconv.Atoi(topicId)
	if err != nil || topic <= 0 {
		log.Printf("Parâmetro de página inválido: %v, usando 1 como padrão", topicId)
		page = 1
	}
	matter, err := strconv.Atoi(matterId)
	if err != nil || matter <= 0 {
		log.Printf("Parâmetro de página inválido: %v, usando 1 como padrão", matterId)
		page = 1
	}
	university, err := strconv.Atoi(universityId)
	if err != nil || matter <= 0 {
		log.Printf("Parâmetro de página inválido: %v, usando 1 como padrão", universityId)
		page = 1
	}

	offset := page - 1
	limit := 10

	if topic == 0 || dateInitial == "" || dateFinal == "" || university == 0 {
		http.Error(w, "Params is required", http.StatusBadRequest)
		return
	}

	filterQuestions := models.FilterQuestions{
		Topic:       topic,
		Matter:      matter,
		University:  university,
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
