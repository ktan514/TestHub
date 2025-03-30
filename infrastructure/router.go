// ====================================
// ファイル名: router.go
// 作成日: 2025/03/27
// 作成者: K.Takeuchi
// 説明: このファイルの概要を記述
// ====================================

package infrastructure

import (
	"testhub/internal/interface/controller"
	"testhub/internal/interface/gateway"
	"testhub/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {

	e := echo.New()

	// ログを出すミドルウェア
	e.Use(middleware.Logger())

	// リカバリ（panicが起きたら500で返す）
	e.Use(middleware.Recover())

	// CORSを許可
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	// カテゴリルーティング
	{
		repo := gateway.NewTestCategoryRepository()
		u := usecase.NewTestCategoryUsecase(repo)
		testCategoryController := controller.NewTestCategoryController(u)
		e.GET("/api/categories", testCategoryController.GetTestCategories)
	}

	// ユーザールーティング
	{
		repo := gateway.NewUserRepository()
		u := usecase.NewUserUsecase(repo)
		userController := controller.NewUserController(u)
		e.GET("/api/users", userController.GetUsers)
		e.POST("/api/login", userController.Login)
	}

	// テスト項目ルーティング
	{
		repo := gateway.NewTestItemRepository()
		u := usecase.NewTestItemUsecase(repo)
		testItemController := controller.NewTestItemController(u)
		e.GET("/api/testitems", testItemController.GetTestItems)
		e.DELETE("/api/testitems", testItemController.DelTestItems)

		e.POST("/api/testitem", testItemController.AddTestItem)
		e.GET("/api/testitem", testItemController.GetTestItem)
	}

	return e
}
