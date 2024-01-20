package controller

import (
	"fmt"
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"net/url"
	"os"
	"log"
	"time"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
	GetCookies(c echo.Context) error 
	GetUser(c echo.Context) error
	GetToken(c echo.Context) error 
	AuthLogin(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}

func (uc *userController) LogIn(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, userRes, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	// JWTトークンをCookieに保存
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(3 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)

	// ユーザー情報をCookieに保存
	userInfo, err := json.Marshal(userRes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	encodedUserInfo := url.QueryEscape(string(userInfo))
	userInfoCookie := new(http.Cookie)
	userInfoCookie.Name = "userInfo"
	userInfoCookie.Value = encodedUserInfo
	userInfoCookie.Expires = time.Now().Add(3 * time.Hour)
	userInfoCookie.Path = "/"
	userInfoCookie.Domain = os.Getenv("API_DOMAIN")
	userInfoCookie.Secure = true
	userInfoCookie.HttpOnly = true
	userInfoCookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(userInfoCookie)

	// JWTトークンとユーザー情報をレスポンスボディに含める
	return c.JSON(http.StatusOK, echo.Map{
		"token": tokenString,
		"user":  userRes,
	})
}


func (uc *userController) LogOut(c echo.Context) error {
	// トークン用のクッキーを削除
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)

	// ユーザー情報用のクッキーを削除
	userCookie := new(http.Cookie)
	userCookie.Name = "userInfo"
	userCookie.Value = ""
	userCookie.Expires = time.Now() 
	userCookie.Path = "/"
	userCookie.Domain = os.Getenv("API_DOMAIN")
	userCookie.Secure = true
	userCookie.HttpOnly = true
	userCookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(userCookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}

func (uc *userController) GetCookies(c echo.Context) error {
    // リクエストから全てのクッキーを取得
    cookies := c.Cookies()

    // クッキーの情報を格納するためのマップ
    cookieMap := make(map[string]string)
    for _, cookie := range cookies {
			fmt.Println("Cookie:", cookie.Name, "Value:", cookie.Value) // この行を追加
        cookieMap[cookie.Name] = cookie.Value
    }

    // クッキーの情報をJSONとして返す
    return c.JSON(http.StatusOK, cookieMap)
}

// userControllerに追加するGetUserメソッド
func (uc *userController) GetUser(c echo.Context) error {
    // JWTトークンからユーザーIDを取得
    userToken := c.Get("user").(*jwt.Token)
    claims := userToken.Claims.(jwt.MapClaims)
    
    // userIDをfloat64として取得し、その後uintに変換
    userIDFloat, ok := claims["user_id"].(float64)
    if !ok {
        return c.JSON(http.StatusBadRequest, "Invalid user ID format")
    }
    userID := uint(userIDFloat)

    // データベースからユーザー情報を取得
    userInfo, err := uc.uu.GetUserByID(userID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, userInfo)
}

func (uc *userController) GetToken(c echo.Context) error {
    // トークン自体を取得（トークンはリクエストのクッキーに保存されている）
    token, err := c.Cookie("token")
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "Token retrieval failed")
    }

    // トークンをレスポンスとして返す
    return c.JSON(http.StatusOK, map[string]string{"token": token.Value})
}

// userControllerに追加するAuthLoginメソッド
func (uc *userController) AuthLogin(c echo.Context) error {
    fmt.Println("AuthLogin method called")
    var user model.User
    if err := c.Bind(&user); err != nil {
        log.Printf("Error in AuthLogin: %v", err)
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
    }

    authenticatedUser, err := uc.uu.AuthenticateUser(user)
    if err != nil {
        log.Printf("Authentication failed: %v", err)
        return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Authentication failed"})
    }

    // JWTトークンの生成
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": authenticatedUser.ID,
        "name":  authenticatedUser.Name,
        "email": authenticatedUser.Email,
        "exp":   time.Now().Add(time.Hour * 24).Unix(), // トークンの有効期限
    })

    // 環境変数からJWTシークレットを取得
    secretKey := os.Getenv("SECRET")
    if secretKey == "" {
        log.Printf("JWT secret key is not set")
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
    }

    // トークンに署名
    tokenString, err := token.SignedString([]byte(secretKey))
    if err != nil {
        log.Printf("Token signing error: %v", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Token signing error"})
    }

    // JWTトークンを含むレスポンスを返す
    return c.JSON(http.StatusOK, map[string]interface{}{
        "user": authenticatedUser,
        "jwt":  tokenString,
    })
}
