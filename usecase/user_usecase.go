package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
	"os"
	"log"
	"time"
	"errors" 
  "gorm.io/gorm"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, model.UserResponse, error)
	GetUserByID(userID uint) (model.UserResponse, error)
	AuthenticateUser(user model.User) (model.UserResponse, error)
	FindOrCreateUser(email, name string) (model.UserResponse, error)
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{
		Email: user.Email,
		Password: string(hash),
		Name:     user.Name,
	}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
		Name:  newUser.Name,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, model.UserResponse, error) {
	// ログインバリデーション関数を呼び出す（UserLoginValidateは新しく作成する必要があります）
	if err := uu.uv.UserLoginValidate(user); err != nil {
		return "", model.UserResponse{}, err
	}
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", model.UserResponse{}, err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", model.UserResponse{}, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:    storedUser.ID,
		Email: storedUser.Email,
		Name:  storedUser.Name,
	}
	return tokenString, resUser, nil
}

// GetUserByID メソッドを追加
func (uu *userUsecase) GetUserByID(userID uint) (model.UserResponse, error) {
    user := model.User{}
    err := uu.ur.GetUserById(&user, userID)
    if err != nil {
        return model.UserResponse{}, err
    }
    return model.UserResponse{
        ID:    user.ID,
        Email: user.Email,
        Name:  user.Name,
    }, nil
}

func (uu *userUsecase) AuthenticateUser(user model.User) (model.UserResponse, error) {
    log.Println("AuthenticateUser method called") // この行を追加

    // ユーザーが提供したメールアドレスでDBを検索
    storedUser := model.User{}
    log.Println("Searching user by email") // この行を追加
    if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
        log.Printf("User not found: %v", err) // この行を追加
        return model.UserResponse{}, err
    }

    // パスワードの検証
    log.Println("Validating password") // この行を追加
    if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
        log.Printf("Password validation failed: %v", err) // この行を追加
        return model.UserResponse{}, err
    }

    log.Println("User authenticated successfully") // この行を追加

    // 認証が成功した場合、ユーザー情報を返す
    return model.UserResponse{
        ID:    storedUser.ID,
        Email: storedUser.Email,
        Name:  storedUser.Name,
    }, nil
}

func (uu *userUsecase) FindOrCreateUser(email, name string) (model.UserResponse, error) {
    var user model.User
    err := uu.ur.GetUserByEmail(&user, email)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            // ユーザーが存在しない場合、新規作成
            user = model.User{
                Email: email,
                Name:  name,
                // パスワードはOAuthの場合、空またはランダムに生成した値を設定
            }
            if err := uu.ur.CreateUser(&user); err != nil {
                return model.UserResponse{}, err
            }
        } else {
            // データベースエラー
            return model.UserResponse{}, err
        }
    }

    // ユーザーが既に存在するか、新規作成された場合のレスポンス
    return model.UserResponse{
        ID:    user.ID,
        Email: user.Email,
        Name:  user.Name,
    }, nil
}
