package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	// Rotas da API
	router.HandleFunc("/api/lessons", GetLessons).Methods("GET")
	router.HandleFunc("/api/lessons/{id}", GetLesson).Methods("GET")
	router.HandleFunc("/api/user/progress", UpdateUserProgress).Methods("POST")
	router.HandleFunc("/api/user/login", Login).Methods("POST")
	router.HandleFunc("/api/user/register", Register).Methods("POST")

	// Configuração CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func GetLessons(w http.ResponseWriter, r *http.Request) {
	// Lógica para buscar lições
}

func GetLesson(w http.ResponseWriter, r *http.Request) {
	// Lógica para buscar uma lição específica
}

func UpdateUserProgress(w http.ResponseWriter, r *http.Request) {
	// Lógica para atualizar progresso do usuário
}

func Login(w http.ResponseWriter, r *http.Request) {
	// Lógica de autenticação
}

func Register(w http.ResponseWriter, r *http.Request) {
	// Lógica de registro
}
