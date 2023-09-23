package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IShopController interface {
	GetAllShops(c echo.Context) error
	GetShopById(c echo.Context) error
	CreateShop(c echo.Context) error
	UpdateShop(c echo.Context) error
	DeleteShop(c echo.Context) error
}

type shopController struct {
	su usecase.IShopUsecase
}

func NewShopController(su usecase.IShopUsecase) IShopController {
	return &shopController{su}
}

func (sc *shopController) GetAllShops(c echo.Context) error {
	shopsRes, err := sc.su.GetAllShops()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, shopsRes)
}

func (sc *shopController) GetShopById(c echo.Context) error {
	id := c.Param("shopId")
	shopId, _ := strconv.Atoi(id)
	shopRes, err := sc.su.GetShopById(uint(shopId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, shopRes)
}

func (sc *shopController) CreateShop(c echo.Context) error {
	shop := model.Shop{}
	if err := c.Bind(&shop); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	shopRes, err := sc.su.CreateShop(shop)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, shopRes)
}

func (sc *shopController) UpdateShop(c echo.Context) error {
	id := c.Param("shopId")
	shopId, _ := strconv.Atoi(id)

	shop := model.Shop{}
	if err := c.Bind(&shop); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	shopRes, err := sc.su.UpdateShop(shop, uint(shopId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, shopRes)
}

func (sc *shopController) DeleteShop(c echo.Context) error {
	id := c.Param("shopId")
	shopId, _ := strconv.Atoi(id)

	err := sc.su.DeleteShop(uint(shopId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

