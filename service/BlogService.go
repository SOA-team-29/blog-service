package service

import (
	"blog/model"
	"blog/repo"
)

type BlogService struct {
	BlogRepo *repo.BlogRepository
}

func (service *BlogService) CreateBlog(blog *model.Blog) error {
	err := service.BlogRepo.CreateBlog(blog)
	if err != nil {
		return err
	}
	return nil
}
