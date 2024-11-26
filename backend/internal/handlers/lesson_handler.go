package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"codequest/internal/models"
	"github.com/gorilla/mux"
)

var lessons = []models.Lesson{
	{ID: 1, Title: "Introdução ao Python", Description: "Aprenda os conceitos básicos de Python", Language: "Python", Difficulty: "Iniciante"},
	{ID: 2, Title: "Funções em JavaScript", Description: "Como criar e usar funções em JavaScript", Language: "JavaScript", Difficulty: "Intermediário"},
}

func GetLessons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lessons)
}

func GetLessonByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for _, lesson := range lessons {
		if lesson.ID == id {
			json.NewEncoder(w).Encode(lesson)
			return
		}
	}

	http.Error(w, "Lição não encontrada", http.StatusNotFound)
}
