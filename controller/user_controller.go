package controller

import (
	"fmt"
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"net/url"
	"os"
	"time"
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
	GetCookies(c echo.Context) error 
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
	// トークンをCookieに保存
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
	userCookie := new(http.Cookie)
	userCookie.Name = "userInfo"
	userCookie.Value = encodedUserInfo
	userCookie.Expires = time.Now().Add(3 * time.Hour)
	userCookie.Path = "/"
	userCookie.Domain = "https://ecsite-lqh9.onrender.com"
	userCookie.Secure = true
	userCookie.HttpOnly = true
	userCookie.SameSite = http.SameSiteNoneMode
	fmt.Println("Encoded User Cookie Value:", userCookie.Value)
	c.SetCookie(userCookie)
	return c.NoContent(http.StatusOK)
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
