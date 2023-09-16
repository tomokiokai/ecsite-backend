package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IBlogRepository interface {
	GetAllBlogs(blogs *[]model.Blog, userId uint) error
	GetBlogById(blog *model.Blog, userId uint, blogId uint) error
	CreateBlog(blog *model.Blog) error
	UpdateBlog(blog *model.Blog, userId uint, blogId uint) error
	DeleteBlog(userId uint, blogId uint) error
	GetAllBlogsForBuild() ([]model.Blog, error)
}

type blogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) IBlogRepository {
	return &blogRepository{db}
}

func (br *blogRepository) GetAllBlogs(blogs *[]model.Blog, userId uint) error {
	if err := br.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(blogs).Error; err != nil {
		return err
	}
	return nil
}

func (br *blogRepository) GetBlogById(blog *model.Blog, userId uint, blogId uint) error {
	if err := br.db.Joins("User").Where("user_id=?", userId).First(blog, blogId).Error; err != nil {
		return err
	}
	return nil
}

func (br *blogRepository) CreateBlog(blog *model.Blog) error {
	if err := br.db.Create(blog).Error; err != nil {
		return err
	}
	return nil
}

func (br *blogRepository) UpdateBlog(blog *model.Blog, userId uint, blogId uint) error {
	result := br.db.Model(blog).Clauses(clause.Returning{}).Where("id=? AND user_id=?", blogId, userId).Updates(map[string]interface{}{
		"title":   blog.Title,
		"content": blog.Content,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (br *blogRepository) DeleteBlog(userId uint, blogId uint) error {
	result := br.db.Where("id=? AND user_id=?", blogId, userId).Delete(&model.Blog{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (br *blogRepository) GetAllBlogsForBuild() ([]model.Blog, error) {
    var blogs []model.Blog
    if err := br.db.Joins("User").Order("created_at").Find(&blogs).Error; err != nil {
        return nil, err
    }
    return blogs, nil
}

