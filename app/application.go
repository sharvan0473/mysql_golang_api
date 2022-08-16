package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	fmt.Println("Server Started")
	mapUrls()

	http.ListenAndServe(":8000", router)
}
