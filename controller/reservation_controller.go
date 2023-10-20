package controller

import (
	"net/http"
	"go-rest-api/model"
	"go-rest-api/usecase"
	"github.com/labstack/echo/v4"
	"strconv" 
    "github.com/golang-jwt/jwt/v4"
)

type IReservationController interface {
    MakeReservation(c echo.Context) error
    CancelReservation(c echo.Context) error
    GetReservationByUser(c echo.Context) error
    GetAllReservations(c echo.Context) error
    UpdateReservation(c echo.Context) error
	GetReservationsForBuild(c echo.Context) error
}

type reservationController struct {
    ru usecase.IReservationUsecase
}

func NewReservationController(ru usecase.IReservationUsecase) IReservationController {
    return &reservationController{ru}
}

func (rc *reservationController) MakeReservation(c echo.Context) error {
    // URLからshopIdを取得
    shopIdParam := c.Param("shopId")

    // shopIdが空でないことを確認します。
    if shopIdParam == "" {
        return c.JSON(http.StatusBadRequest, "Shop ID is required")
    }

    // shopIdを整数に変換します。エラーがあれば処理します。
    shopId, err := strconv.Atoi(shopIdParam)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Shop ID must be an integer")
    }

    // JWTトークンからユーザー情報を取得
    user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)

    // claims["user_id"] は float64 型であるため、型アサーションを適切に行います。
    userIdFloat, ok := claims["user_id"].(float64) // 修正: 型アサーションをfloat64に
    if !ok {
        return c.JSON(http.StatusInternalServerError, "Invalid user ID")
    }

    // float64 から uint への変換
    userId := uint(userIdFloat) // 修正: float64をuintに変換

    // リクエストボディから予約データを取得
    reservation := model.Reservation{}
    if err := c.Bind(&reservation); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    // ユーザーIDとショップIDをreservationモデルに設定
    reservation.UserID = userId // ここでユーザーIDを設定します。
    reservation.ShopID = uint(shopId)

    // 予約を作成
    reservationRes, err := rc.ru.MakeReservation(reservation)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusCreated, reservationRes)
}

func (rc *reservationController) CancelReservation(c echo.Context) error {
    reservationId := c.Param("reservationId")
    err := rc.ru.CancelReservation(reservationId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.NoContent(http.StatusNoContent)
}

func (rc *reservationController) GetReservationByUser(c echo.Context) error {
    userId := c.Param("userId")
    if userId == "" {
        return c.JSON(http.StatusBadRequest, "User ID is required")
    }
    reservationsRes, err := rc.ru.GetReservationByUser(userId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, reservationsRes)
}

func (rc *reservationController) GetAllReservations(c echo.Context) error {
    reservationsRes, err := rc.ru.GetAllReservations()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, reservationsRes)
}

func (rc *reservationController) UpdateReservation(c echo.Context) error {
    reservation := model.Reservation{}
    if err := c.Bind(&reservation); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    updatedReservation, err := rc.ru.UpdateReservation(reservation)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, updatedReservation)
}

func (rc *reservationController) GetReservationsForBuild(c echo.Context) error {
    // ビルドプロセス用に特別に設計されたロジックで予約情報を取得
    reservations, err := rc.ru.GetReservationsForBuild()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, reservations)
}
