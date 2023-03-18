package main

import (
	"PostApp/handlers"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost:8000
	}

	// Posts routes
	router.HandleFunc("/api/posts", handlers.GetAllPosts).Methods("GET")
	router.HandleFunc("/api/posts/{id}", handlers.GetPostByID).Methods("GET")
	router.HandleFunc("/api/posts", handlers.CreateNewPost).Methods("POST")
	router.HandleFunc("/api/posts/{id}", handlers.DeletePost).Methods("DELETE")
	router.HandleFunc("/api/posts/{id}", handlers.UpdatePost).Methods("PATCH")

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, corsMiddleware(router)) // Uygulamamız localhost:8000/api altında istekleri dinlemeye başlar
	if err != nil {
		fmt.Print(err)
	}
}
