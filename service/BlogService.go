package service

import (
	"blog/model"
	"blog/repo"
	"fmt"
)

type BlogService struct {
	BlogRepository *repo.BlogRepository
}

func (service *BlogService) CreateBlog(blog *model.Blog) error {
	err := service.BlogRepository.CreateBlog(blog)
	if err != nil {
		return err
	}
	return nil
}
func (service *BlogService) FindBlog(id string) (*model.Blog, error) {
	blog, err := service.BlogRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("blog with id %s not found", id))
	}
	return &blog, nil
}
func (service *BlogService) GetAll(page, pageSize int) (*[]model.Blog, error) {
	blogs, err := service.BlogRepository.GetAll(page, pageSize)
	if err != nil {
		return nil, err
	}
	return &blogs, nil
}
func (service *BlogService) SetBlogComments(blogID int, comments []model.BlogComment) error {
	err := service.BlogRepository.SetBlogComments(blogID, comments)
	if err != nil {
		return err
	}
	return nil
}
func (service *BlogService) GetBlogByID(ID int) (*model.Blog, error) {
	blog, err := service.BlogRepository.GetBlogByID(ID)
	if err != nil {
		return &model.Blog{}, err
	}
	return &blog, nil
}
