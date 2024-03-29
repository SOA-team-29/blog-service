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

func (blogRepo *BlogRepository) FindById(id string) (model.Blog, error) {
	blog := model.Blog{}
	dbResult := blogRepo.DatabaseConnection.First(&blog, "id = ?", id)
	if dbResult.Error != nil {
		return blog, dbResult.Error
	}
	return blog, nil
}

func (blogRepo *BlogRepository) GetAll(page, pageSize int) ([]model.Blog, error) {

	blogs := []model.Blog{}
	//tourCharacteristics := []model.TourCharacteristic{}
	/*
		dbResult := tourRepo.DatabaseConnection.Omit("tour_characteristics").Find(&tours)
		for i := range tours {
			tourRepo.DatabaseConnection.Model(&model.Tour{}).Where("id=?", tours[i].ID).Pluck("tour_characteristics", &tourCharacteristics)
			tours[i].TourCharacteristics = tourCharacteristics
		}*/
	dbResult := blogRepo.DatabaseConnection.Find(&blogs)
	if dbResult != nil {
		return blogs, dbResult.Error
	}
	return blogs, nil
}
