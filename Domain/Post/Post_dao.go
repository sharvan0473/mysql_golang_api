package domain

import (
	"fmt"
	"math/rand"
	"sharvan/mux/api/data_source/db"
	"sharvan/mux/api/utils/errors"
)

var (
	postDB = make(map[int]*Post)
)

const (
	queryReadAllPost = "SELECT * from posts;"
)

func GetAllPost() ([]Post, *errors.RestErr) {
	results := make([]Post, 0)

	postResults, err := db.ProjectDB.Query(queryReadAllPost)
	if err != nil {
		panic(err.Error())
		return nil, errors.NewInternalServerError("Error in listing")
	}
	defer postResults.Close()

	for postResults.Next() {
		var mPost Post
		err := postResults.Scan(&mPost.ID, &mPost.Title, &mPost.Body)
		if err != nil {
			panic(err.Error())
			return nil, errors.NewInternalServerError("Error in listing ids")
		}
		results = append(results, mPost)
	}

	return results, nil
}

func (post *Post) SavePost() (*Post, *errors.RestErr) {
	mPost := &Post{
		ID:    rand.Intn(100),
		Title: post.Title,
		Body:  post.Body,
	}
	fmt.Println("not Herr")
	if err := post.Validate(); err != nil {
		fmt.Println(" Herr")
		//fmt.Println(err)
		return nil, err
	}
	fmt.Println("mPost Herr")
	postDB[mPost.ID] = mPost
	fmt.Println("mPost Herry===")
	return mPost, nil
}

func (post *Post) UpdatePost() (*Post, *errors.RestErr) {
	mPost := &Post{
		ID:    post.ID,
		Title: post.Title,
		Body:  post.Body,
	}

	mId := post.ID
	_, ok := postDB[mId]

	if !ok {

		return nil, errors.PostNotFoundError("ID No Present in DB")
	}

	if err := post.Validate(); err != nil {
		fmt.Println(" Herr")
		return nil, err
	}

	postDB[mPost.ID] = mPost
	return mPost, nil
}

func GetDetailPost(searchId int) (*Post, *errors.RestErr) {
	_, ok := postDB[searchId]
	if ok {
		return postDB[searchId], nil
	}
	return nil, errors.PostNotFoundError("Not Found")
}
