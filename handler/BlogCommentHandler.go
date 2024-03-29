package handler

import (
	"blog/model"
	"blog/service"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type BlogCommentHandler struct {
	BlogCommentService *service.BlogCommentService
}

func (handler *BlogCommentHandler) GetByBlogId(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["blogId"]
	log.Printf("Searching for comments for blog ID: %s", id)

	comments, err := handler.BlogCommentService.GetByBlogId(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(comments)
}

func (handler *BlogCommentHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var comment model.BlogComment
	err := json.NewDecoder(req.Body).Decode(&comment)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.BlogCommentService.Create(&comment)
	if err != nil {
		println("Error while creating a new comment")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *BlogCommentHandler) Delete(writer http.ResponseWriter, req *http.Request) {

	var comment model.BlogComment
	err := json.NewDecoder(req.Body).Decode(&comment)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	/*
		log.Printf("Deleting comment from blog ID: %s\tcreated @ %s", comment.ID, comment.PublishedDateTime)

		err = handler.BlogCommentService.Delete(comment.ID, comment.PublishedDateTime)
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		writer.WriteHeader(http.StatusOK)*/
}
