package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

type IBlogUsecase interface {
	GetAllBlogs(userId uint) ([]model.BlogResponse, error)
	GetBlogById(userId uint, blogId uint) (model.BlogResponse, error)
	CreateBlog(blog model.Blog) (model.BlogResponse, error)
	UpdateBlog(blog model.Blog, userId uint, blogId uint) (model.BlogResponse, error)
	DeleteBlog(userId uint, blogId uint) error
}

type blogUsecase struct {
	br repository.IBlogRepository
	bv validator.IBlogValidator
}

func NewBlogUsecase(br repository.IBlogRepository, bv validator.IBlogValidator) IBlogUsecase {
	return &blogUsecase{br, bv}
}

func (bu *blogUsecase) GetAllBlogs(userId uint) ([]model.BlogResponse, error) {
	blogs := []model.Blog{}
	if err := bu.br.GetAllBlogs(&blogs, userId); err != nil {
		return nil, err
	}
	resBlogs := []model.BlogResponse{}
	for _, v := range blogs {
		b := model.BlogResponse{
			ID:        v.ID,
			Title:     v.Title,
			Content:   v.Content,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resBlogs = append(resBlogs, b)
	}
	return resBlogs, nil
}

func (bu *blogUsecase) GetBlogById(userId uint, blogId uint) (model.BlogResponse, error) {
	blog := model.Blog{}
	if err := bu.br.GetBlogById(&blog, userId, blogId); err != nil {
		return model.BlogResponse{}, err
	}
	resBlog := model.BlogResponse{
		ID:        blog.ID,
		Title:     blog.Title,
		Content:   blog.Content,
		CreatedAt: blog.CreatedAt,
		UpdatedAt: blog.UpdatedAt,
	}
	return resBlog, nil
}

func (bu *blogUsecase) CreateBlog(blog model.Blog) (model.BlogResponse, error) {
	if err := bu.bv.BlogValidate(blog); err != nil {
		return model.BlogResponse{}, err
	}
	if err := bu.br.CreateBlog(&blog); err != nil {
		return model.BlogResponse{}, err
	}
	resBlog := model.BlogResponse{
		ID:        blog.ID,
		Title:     blog.Title,
		Content:   blog.Content,
		CreatedAt: blog.CreatedAt,
		UpdatedAt: blog.UpdatedAt,
	}
	return resBlog, nil
}

func (bu *blogUsecase) UpdateBlog(blog model.Blog, userId uint, blogId uint) (model.BlogResponse, error) {
	if err := bu.bv.BlogValidate(blog); err != nil {
		return model.BlogResponse{}, err
	}
	if err := bu.br.UpdateBlog(&blog, userId, blogId); err != nil {
		return model.BlogResponse{}, err
	}
	resBlog := model.BlogResponse{
		ID:        blog.ID,
		Title:     blog.Title,
		Content:   blog.Content,
		CreatedAt: blog.CreatedAt,
		UpdatedAt: blog.UpdatedAt,
	}
	return resBlog, nil
}

func (bu *blogUsecase) DeleteBlog(userId uint, blogId uint) error {
	if err := bu.br.DeleteBlog(userId, blogId); err != nil {
		return err
	}
	return nil
}
