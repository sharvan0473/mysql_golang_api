package services

import (
	"fmt"
	domain "sharvan/mux/api/Domain/Post"
	"sharvan/mux/api/utils/errors"
)

func GetAllPostService() ([]domain.Post, *errors.RestErr) {
	results, err := domain.GetAllPost()
	if err != nil {
		return nil, errors.NewInternalServerError("Something Went Wrong")
	}
	return results, nil
}

func CreatePostService(post domain.Post) (*domain.Post, *errors.RestErr) {
	results, err := post.SavePost()
	if err != nil {
		fmt.Println("Saving Error")
		return nil, errors.NewInternalServerError(err.Message)
	}
	return results, nil
}

func UpdatePostService(post *domain.Post) (*domain.Post, *errors.RestErr) {
	result, err := post.UpdatePost()
	if err != nil {
		return nil, errors.PostNotFoundError(err.Message)
	}
	return result, nil
}

func GetPostByIdService(id int) (*domain.Post, *errors.RestErr) {
	result, err := domain.GetDetailPost(id)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Message)
	}
	return result, nil
}
