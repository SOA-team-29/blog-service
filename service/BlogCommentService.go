package service

import (
	"blog/model"
	"blog/repo"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type BlogCommentService struct {
	BlogCommentRepository *repo.BlogCommentRepository
}

func (service *BlogCommentService) FindBlogComment(id string) (*model.BlogComment, error) {
	blogComment, err := service.BlogCommentRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("blog comment with id %s not found", id))
	}
	return &blogComment, nil
}

func (service *BlogCommentService) Create(blogComment *model.BlogComment) error {
	err := service.BlogCommentRepository.Create(blogComment)
	if err != nil {
		return err
	}
	return nil
}

func (service *BlogCommentService) GetAll() ([]model.BlogComment, error) {
	blogComments, err := service.BlogCommentRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return blogComments, nil
}

func (service *BlogCommentService) GetByBlogId(blogId string) ([]model.BlogComment, error) {
	blogComments, err := service.BlogCommentRepository.GetByBlogId(blogId)
	if err != nil {
		return nil, err
	}
	return blogComments, nil
}

func (service *BlogCommentService) Delete(blogId uuid.UUID, commentCreateTime time.Time) error {
	err := service.BlogCommentRepository.Delete(blogId, commentCreateTime)
	if err != nil {
		return err
	}
	return nil
}
