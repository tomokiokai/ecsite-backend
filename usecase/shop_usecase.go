package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

type IShopUsecase interface {
	GetAllShops() ([]model.ShopResponse, error)
	GetShopById(shopId uint) (model.ShopResponse, error)
	CreateShop(shop model.Shop) (model.ShopResponse, error)
	UpdateShop(shop model.Shop, shopId uint) (model.ShopResponse, error)
	DeleteShop(shopId uint) error
}

type shopUsecase struct {
	sr repository.IShopRepository
	sv validator.IShopValidator
}

func NewShopUsecase(sr repository.IShopRepository, sv validator.IShopValidator) IShopUsecase {
	return &shopUsecase{sr, sv}
}

func (su *shopUsecase) GetAllShops() ([]model.ShopResponse, error) {
	shops := []model.Shop{}
	if err := su.sr.GetAllShops(&shops); err != nil {
		return nil, err
	}
	resShops := []model.ShopResponse{}
	for _, v := range shops {
		s := model.ShopResponse{
			ID:          v.ID,
			Name:        v.Name,
			Address:     v.Address,
			Area:        v.Area,
			Genre:       v.Genre,
			Description: v.Description,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		}
		resShops = append(resShops, s)
	}
	return resShops, nil
}

func (su *shopUsecase) GetShopById(shopId uint) (model.ShopResponse, error) {
	shop := model.Shop{}
	if err := su.sr.GetShopById(&shop, shopId); err != nil {
		return model.ShopResponse{}, err
	}
	resShop := model.ShopResponse{
		ID:          shop.ID,
		Name:        shop.Name,
		Address:     shop.Address,
		Area:        shop.Area,
		Genre:       shop.Genre,
		Description: shop.Description,
		CreatedAt:   shop.CreatedAt,
		UpdatedAt:   shop.UpdatedAt,
	}
	return resShop, nil
}

func (su *shopUsecase) CreateShop(shop model.Shop) (model.ShopResponse, error) {
	if err := su.sv.ShopValidate(shop); err != nil {
		return model.ShopResponse{}, err
	}
	if err := su.sr.CreateShop(&shop); err != nil {
		return model.ShopResponse{}, err
	}
	resShop := model.ShopResponse{
		ID:          shop.ID,
		Name:        shop.Name,
		Address:     shop.Address,
		Area:        shop.Area,
		Genre:       shop.Genre,
		Description: shop.Description,
		CreatedAt:   shop.CreatedAt,
		UpdatedAt:   shop.UpdatedAt,
	}
	return resShop, nil
}

func (su *shopUsecase) UpdateShop(shop model.Shop, shopId uint) (model.ShopResponse, error) {
	if err := su.sv.ShopValidate(shop); err != nil {
		return model.ShopResponse{}, err
	}
	if err := su.sr.UpdateShop(&shop, shopId); err != nil {
		return model.ShopResponse{}, err
	}
	resShop := model.ShopResponse{
		ID:          shop.ID,
		Name:        shop.Name,
		Address:     shop.Address,
		Area:        shop.Area,
		Genre:       shop.Genre,
		Description: shop.Description,
		CreatedAt:   shop.CreatedAt,
		UpdatedAt:   shop.UpdatedAt,
	}
	return resShop, nil
}

func (su *shopUsecase) DeleteShop(shopId uint) error {
	if err := su.sr.DeleteShop(shopId); err != nil {
		return err
	}
	return nil
}
