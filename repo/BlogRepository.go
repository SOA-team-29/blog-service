package repo

import (
	"blog/model"

	"gorm.io/gorm"
)

type BlogRepository struct {
	DatabaseConnection *gorm.DB
}

func (blogRepo *BlogRepository) CreateBlog(blog *model.Blog) error {
	dbResult := blogRepo.DatabaseConnection.Create(blog)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
