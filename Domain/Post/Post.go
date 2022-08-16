package domain

import "sharvan/mux/api/utils/errors"

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (post *Post) Validate() *errors.RestErr {
	if post.Title == "" {
		return errors.NewBadRequest("Empty Post Title")
	}
	if post.Body == "" {
		return errors.NewBadRequest("Empty Post Body")
	}
	return nil
}
