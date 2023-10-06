package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

type IFavoriteUsecase interface {
	AddFavorite(favorite model.Favorite) (model.FavoriteResponse, error)
	RemoveFavorite(fshopId, userId string) error
	GetFavorites(userId string) ([]model.FavoriteResponse, error)
	GetFavoriteShops(userId string) ([]model.ShopResponse, error)
	GetFavoritesForBuild() ([]model.FavoriteResponse, error)
}

type favoriteUsecase struct {
	fr repository.IFavoriteRepository
	sr repository.IShopRepository  
	ur repository.IUserRepository  
	fv validator.IFavoriteValidator 
}

func NewFavoriteUsecase(fr repository.IFavoriteRepository, sr repository.IShopRepository, ur repository.IUserRepository, fv validator.IFavoriteValidator) IFavoriteUsecase {
	return &favoriteUsecase{fr, sr, ur, fv}
}

func (fu *favoriteUsecase) AddFavorite(favorite model.Favorite) (model.FavoriteResponse, error) {
	if err := fu.fv.FavoriteValidate(favorite); err != nil {
		return model.FavoriteResponse{}, err
	}

	// ここで repository の AddFavorite メソッドを呼び出す
	err := fu.fr.AddFavorite(&favorite)
	if err != nil {
		return model.FavoriteResponse{}, err
	}

	shop := model.Shop{}
	err = fu.sr.GetShopById(&shop, favorite.ShopID)
	if err != nil {
		return model.FavoriteResponse{}, err
	}

	user := model.User{}
	err = fu.ur.GetUserById(&user, favorite.UserID)
	if err != nil {
		return model.FavoriteResponse{}, err
	}

	return model.FavoriteResponse{
		ID:     favorite.ID,
		Shop:   shop,
		User:   user,
	}, nil
}

func (fu *favoriteUsecase) RemoveFavorite(shopId, userId string) error {
    return fu.fr.RemoveFavorite(shopId, userId)
}


func (fu *favoriteUsecase) GetFavorites(userId string) ([]model.FavoriteResponse, error) {
	favorites := []model.Favorite{}
	if err := fu.fr.GetFavorites(userId, &favorites); err != nil {
		return nil, err
	}
	resFavorites := []model.FavoriteResponse{}
	for _, v := range favorites {
		shop := model.Shop{}
		err := fu.sr.GetShopById(&shop, v.ShopID)
		if err != nil {
			return nil, err
		}
		user := model.User{}
		err = fu.ur.GetUserById(&user, v.UserID)
		if err != nil {
			return nil, err
		}
		resFavorites = append(resFavorites, model.FavoriteResponse{
			ID:     v.ID,
			Shop:   shop,
			User:   user,
			IsFavorite: true,
		})
	}
	return resFavorites, nil
}

func (fu *favoriteUsecase) GetFavoriteShops(userId string) ([]model.ShopResponse, error) {
	shops := []model.Shop{}
	if err := fu.fr.GetFavoriteShops(userId, &shops); err != nil {
		return nil, err
	}
	resShops := []model.ShopResponse{}
	for _, v := range shops {
		resShops = append(resShops, model.ShopResponse{
			ID:          v.ID,
			Name:        v.Name,
			Address:     v.Address,
			Area:        v.Area,
			Genre:       v.Genre,
			Description: v.Description,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}
	return resShops, nil
}

func (fu *favoriteUsecase) GetFavoritesForBuild() ([]model.FavoriteResponse, error) {
	favorites := []model.Favorite{}
	if err := fu.fr.GetFavoritesForBuild(&favorites); err != nil {
		return nil, err
	}
	resFavorites := []model.FavoriteResponse{}
	for _, v := range favorites {
		shop := model.Shop{}
		err := fu.sr.GetShopById(&shop, v.ShopID)
		if err != nil {
			return nil, err
		}
		user := model.User{}
		err = fu.ur.GetUserById(&user, v.UserID)
		if err != nil {
			return nil, err
		}
		resFavorites = append(resFavorites, model.FavoriteResponse{
			ID:     v.ID,
			Shop:   shop,
			User:   user,
			IsFavorite: true,
		})
	}
	return resFavorites, nil
}

