package handler

import (
	"blog/model"
	"blog/service"
	"encoding/json"
	"net/http"
)

type BlogHandler struct {
	BlogService *service.BlogService
}

func (blogHandler *BlogHandler) CreateBlog(writer http.ResponseWriter, req *http.Request) {
	var blog model.Blog
	err := json.NewDecoder(req.Body).Decode(&blog)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = blogHandler.BlogService.CreateBlog(&blog)
	if err != nil {
		println("Error while creating a new tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
