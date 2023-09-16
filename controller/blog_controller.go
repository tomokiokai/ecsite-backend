package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IBlogController interface {
	GetAllBlogs(c echo.Context) error
	GetBlogById(c echo.Context) error
	CreateBlog(c echo.Context) error
	UpdateBlog(c echo.Context) error
	DeleteBlog(c echo.Context) error
	GetBlogsForBuild(c echo.Context) error
}

type blogController struct {
	bu usecase.IBlogUsecase
}

func NewBlogController(bu usecase.IBlogUsecase) IBlogController {
	return &blogController{bu}
}

func (bc *blogController) GetAllBlogs(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	blogsRes, err := bc.bu.GetAllBlogs(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, blogsRes)
}

func (bc *blogController) GetBlogById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("blogId")
	blogId, _ := strconv.Atoi(id)
	blogRes, err := bc.bu.GetBlogById(uint(userId.(float64)), uint(blogId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, blogRes)
}

func (bc *blogController) CreateBlog(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	blog := model.Blog{}
	if err := c.Bind(&blog); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	blog.UserId = uint(userId.(float64))
	blogRes, err := bc.bu.CreateBlog(blog)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, blogRes)
}

func (bc *blogController) UpdateBlog(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("blogId")
	blogId, _ := strconv.Atoi(id)

	blog := model.Blog{}
	if err := c.Bind(&blog); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	blogRes, err := bc.bu.UpdateBlog(blog, uint(userId.(float64)), uint(blogId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, blogRes)
}

func (bc *blogController) DeleteBlog(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("blogId")
	blogId, _ := strconv.Atoi(id)

	err := bc.bu.DeleteBlog(uint(userId.(float64)), uint(blogId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (bc *blogController) GetBlogsForBuild(c echo.Context) error {
    // すべてのユーザーのブログを取得
    blogs, err := bc.bu.GetAllBlogsForBuild()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, blogs)
}

