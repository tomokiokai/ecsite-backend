package router

import (
	"go-rest-api/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// APIキーを検証するカスタムミドルウェア
func ValidateBuildAPIKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiKey := c.Request().Header.Get("X-BUILD-API-KEY")
		if apiKey != os.Getenv("BUILD_API_KEY") {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid build API key")
		}
		return next(c)
	}
}

func NewRouter(
    uc controller.IUserController, 
    tc controller.ITaskController, 
    bc controller.IBlogController, 
    sc controller.IShopController,
    fc controller.IFavoriteController,
    rc controller.IReservationController, 
) *echo.Echo {
	e := echo.New()

	// CORSミドルウェアの設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken,"Authorization" },
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))

	// CSRFミドルウェアの設定
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		TokenLookup:    "header:X-CSRF-Token",
	}))

	// ユーザー関連のエンドポイント
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)

	// CSRFミドルウェアを適用しないエンドポイントのグループ
	s := e.Group("")
	s.GET("/shops", sc.GetAllShops)
	s.GET("/shops/:shopId", sc.GetShopById)

	// tasksエンドポイントの設定
	t := e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteTask)

	// blogsエンドポイントの設定
	b := e.Group("/blogs")
	b.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "header:Authorization",
	}))
	b.GET("", bc.GetAllBlogs)
	b.GET("/:blogId", bc.GetBlogById)
	b.POST("", bc.CreateBlog)
	b.PUT("/:blogId", bc.UpdateBlog)
	b.DELETE("/:blogId", bc.DeleteBlog)

	// お気に入りエンドポイントの設定
    f := e.Group("/favorites")
    f.Use(echojwt.WithConfig(echojwt.Config{
        SigningKey:  []byte(os.Getenv("SECRET")),
        TokenLookup: "header:Authorization",
    }))
    f.POST("", fc.AddFavorite)  // お気に入りを追加
    f.DELETE("/:shopId/:userId", fc.RemoveFavorite)  // お気に入りを削除
    f.GET("", fc.GetFavorites)  // お気に入りを取得

	// reserveエンドポイントの設定
		r := e.Group("/reservations")
    r.Use(echojwt.WithConfig(echojwt.Config{
        SigningKey:  []byte(os.Getenv("SECRET")),
        TokenLookup: "header:Authorization",
    }))
		r.POST("/shop/:shopId", rc.MakeReservation)
    r.DELETE("/:reservationId", rc.CancelReservation)
    r.GET("/user/:userId", rc.GetReservationByUser)
    r.GET("", rc.GetAllReservations)
    r.PUT("/:reservationId", rc.UpdateReservation)

	// ビルド専用のエンドポイント
	build := e.Group("/build")
	build.Use(ValidateBuildAPIKey)  // カスタムミドルウェアを適用
	build.GET("/blogs", bc.GetBlogsForBuild)
	build.GET("/favorites", fc.GetFavoritesForBuild)
	build.GET("/reservations", rc.GetReservationsForBuild)

	return e
}
