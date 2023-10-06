package controller

import (
	"fmt"
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IFavoriteController interface {
	AddFavorite(c echo.Context) error
	RemoveFavorite(c echo.Context) error
	GetFavorites(c echo.Context) error
	GetFavoritesForBuild(c echo.Context) error
	GetFavoriteShops(c echo.Context) error
}

type favoriteController struct {
	fu usecase.IFavoriteUsecase
}

func NewFavoriteController(fu usecase.IFavoriteUsecase) IFavoriteController {
	return &favoriteController{fu}
}

func (fc *favoriteController) AddFavorite(c echo.Context) error {
	favorite := model.Favorite{}
	if err := c.Bind(&favorite); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	favoriteRes, err := fc.fu.AddFavorite(favorite)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, favoriteRes)
}

func (fc *favoriteController) RemoveFavorite(c echo.Context) error {
    shopId := c.Param("shopId")
    userId := c.Param("userId")
    err := fc.fu.RemoveFavorite(shopId, userId)
    if err != nil {
        fmt.Println(err)
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.NoContent(http.StatusNoContent)
}


func (fc *favoriteController) GetFavorites(c echo.Context) error {
	userId := c.Param("userId")  // 仮定: userIdはURLパラメータとして提供されます
	if userId == "" {
		return c.JSON(http.StatusBadRequest, "User ID is required")
	}
	favoritesRes, err := fc.fu.GetFavorites(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, favoritesRes)
}

func (fc *favoriteController) GetFavoritesForBuild(c echo.Context) error {
	favoritesRes, err := fc.fu.GetFavoritesForBuild()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, favoritesRes)
}

func (fc *favoriteController) GetFavoriteShops(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    userID := claims["user_id"].(string)
	favoriteShopsRes, err := fc.fu.GetFavoriteShops(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, favoriteShopsRes)
}