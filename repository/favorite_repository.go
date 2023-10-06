package repository

import (
	"fmt"
	"go-rest-api/model"
	"gorm.io/gorm"
)

type IFavoriteRepository interface {
	AddFavorite(favorite *model.Favorite) error
	RemoveFavorite(shopId, userId string) error
	GetFavorites(userId string, favorites *[]model.Favorite) error
	GetFavoriteShops(userId string, shops *[]model.Shop) error
	GetFavoritesForBuild(favorites *[]model.Favorite) error
}

type favoriteRepository struct {
	db *gorm.DB
}

func NewFavoriteRepository(db *gorm.DB) IFavoriteRepository {
	return &favoriteRepository{db}
}

func (fr *favoriteRepository) AddFavorite(favorite *model.Favorite) error {
	if err := fr.db.Create(favorite).Error; err != nil {
		return err
	}
	return nil
}

func (fr *favoriteRepository) RemoveFavorite(shopId, userId string) error {
    result := fr.db.Where("shop_id=? AND user_id=?", shopId, userId).Delete(&model.Favorite{})
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected < 1 {
        return fmt.Errorf("object does not exist")
    }
    return nil
}


func (fr *favoriteRepository) GetFavorites(userId string, favorites *[]model.Favorite) error {
	if err := fr.db.Where("user_id=?", userId).Find(favorites).Error; err != nil {
		return err
	}
	return nil
}

func (fr *favoriteRepository) GetFavoriteShops(userId string, shops *[]model.Shop) error {
	// SQLクエリを実行してお気に入りのショップを取得します。
	// このクエリは、お気に入りテーブルとショップテーブルを結合し、指定されたユーザーIDに関連するお気に入りのショップを取得します。
	err := fr.db.Joins("JOIN favorites ON favorites.shop_id = shops.id").
		Where("favorites.user_id = ?", userId).
		Find(shops).Error

	// エラーが発生した場合はエラーを返します。
	if err != nil {
		return err
	}

	// 成功した場合はnilを返します。
	return nil
}

func (fr *favoriteRepository) GetFavoritesForBuild(favorites *[]model.Favorite) error {
    if err := fr.db.Preload("Shop").Preload("User").Find(favorites).Error; err != nil {
        return err
    }
    return nil
}