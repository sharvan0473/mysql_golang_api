package app

import "sharvan/mux/api/controllers"

func mapUrls() {
	router.HandleFunc("/posts", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/posts", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", controllers.UpdatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", controllers.GetPostById).Methods("GET")
}
