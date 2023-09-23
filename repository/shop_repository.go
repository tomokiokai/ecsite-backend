package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IShopRepository interface {
	GetAllShops(shops *[]model.Shop) error
	GetShopById(shop *model.Shop, shopId uint) error
	CreateShop(shop *model.Shop) error
	UpdateShop(shop *model.Shop, shopId uint) error
	DeleteShop(shopId uint) error
}

type shopRepository struct {
	db *gorm.DB
}

func NewShopRepository(db *gorm.DB) IShopRepository {
	return &shopRepository{db}
}

func (sr *shopRepository) GetAllShops(shops *[]model.Shop) error {
	if err := sr.db.Order("created_at").Find(shops).Error; err != nil {
		return err
	}
	return nil
}

func (sr *shopRepository) GetShopById(shop *model.Shop, shopId uint) error {
	if err := sr.db.First(shop, shopId).Error; err != nil {
		return err
	}
	return nil
}

func (sr *shopRepository) CreateShop(shop *model.Shop) error {
	if err := sr.db.Create(shop).Error; err != nil {
		return err
	}
	return nil
}

func (sr *shopRepository) UpdateShop(shop *model.Shop, shopId uint) error {
	result := sr.db.Model(shop).Clauses(clause.Returning{}).Where("id=?", shopId).Updates(map[string]interface{}{
		"name":        shop.Name,
		"address":     shop.Address,
		"area":        shop.Area,
		"genre":       shop.Genre,
		"description": shop.Description,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (sr *shopRepository) DeleteShop(shopId uint) error {
	result := sr.db.Where("id=?", shopId).Delete(&model.Shop{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
