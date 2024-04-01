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
	blogComments := []model.BlogComment{}

	dbResult := blogRepo.DatabaseConnection.Omit("blog_comments").Find(&blogs)
	for i := range blogs {
		blogRepo.DatabaseConnection.Model(&model.Blog{}).Where("id=?", blogs[i].ID).Pluck("blog_comments", &blogComments)
		blogs[i].Comments = blogComments
	}
	//dbResult := blogRepo.DatabaseConnection.Find(&blogs)
	if dbResult != nil {
		return blogs, dbResult.Error
	}
	return blogs, nil
}
func (blogRepo *BlogRepository) SetBlogComments(blogID int, comments []model.BlogComment) error {
	var blog model.Blog
	if err := blogRepo.DatabaseConnection.First(&blog, blogID).Error; err != nil {
		return err
	}

	blog.Comments = comments
	if err := blogRepo.DatabaseConnection.Save(&blog).Error; err != nil {
		return err
	}

	return nil
}
func (blogRepo *BlogRepository) GetBlogByID(ID int) (model.Blog, error) {

	blog := model.Blog{}
	blogComments := []model.BlogComment{}

	dbResult := blogRepo.DatabaseConnection.Where("id = ?", ID).Omit("blog_comments").First(&blog)

	blogRepo.DatabaseConnection.Model(&model.Blog{}).Where("id=?", blog.ID).Pluck("blog_comments", &blogComments)
	blog.Comments = blogComments

	if dbResult != nil {
		return blog, dbResult.Error
	}
	return blog, nil
}
