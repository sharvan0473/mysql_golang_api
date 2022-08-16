package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	domain "sharvan/mux/api/Domain/Post"
	"sharvan/mux/api/services"
	"sharvan/mux/api/utils/errors"
	"strconv"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mRes, err := services.GetAllPostService()
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(mRes)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post domain.Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		fmt.Println("Error Here")
		json.NewEncoder(w).Encode(errors.NewBadRequest("Request Body not valid"))
		return
	}
	mResult, saveErr := services.CreatePostService(post)
	if saveErr != nil {
		fmt.Println("saveErr Here")
		json.NewEncoder(w).Encode(errors.NewInternalServerError(saveErr.Message))
		return
	}

	json.NewEncoder(w).Encode(mResult)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post domain.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		json.NewEncoder(w).Encode(errors.NewBadRequest("Request Body Not Valid"))
		return
	}
	params := mux.Vars(r)
	post.ID, err = strconv.Atoi(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(errors.NewBadRequest("Id not Valid"))
		return
	}

	mResult, updateErr := services.UpdatePostService(&post)
	if updateErr != nil {
		json.NewEncoder(w).Encode(errors.PostNotFoundError(updateErr.Message))
		return
	}
	json.NewEncoder(w).Encode(mResult)
}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	searchId, err := strconv.Atoi(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(errors.NewBadRequest("Invalid Id"))
		return
	}
	mResult, getErr := services.GetPostByIdService(searchId)
	if getErr != nil {
		json.NewEncoder(w).Encode(errors.NewBadRequest(getErr.Message))
		return
	}
	json.NewEncoder(w).Encode(mResult)
}
