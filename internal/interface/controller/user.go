// ====================================
// ファイル名: testcategory.go
// 作成日: 2025/03/28
// 作成者: K.Takeuchi
// 説明: このファイルの概要を記述
// ====================================

package controller

import (
	"net/http"
	"testhub/internal/dto"
	"testhub/internal/usecase"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	usecase usecase.UserUsecaseInterface
}

func NewUserController(u usecase.UserUsecaseInterface) *UserController {
	return &UserController{usecase: u}
}

func (ctrl *UserController) GetUsers(c echo.Context) error {
	items, err := ctrl.usecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// DomainからDTO(Data Transfer Object)に変換する
	userDTOs := make([]dto.UserDTO, len(items))
	for i, item := range items {
		userDTOs[i] = dto.UserDTO{
			ID:   item.ID,
			Name: item.Name,
		}
	}

	// JSONに変換して返却する
	return c.JSON(http.StatusOK, userDTOs)
}

// ログイン処理
func (ctrl *UserController) Login(c echo.Context) error {
	// リクエストの構造体
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "リクエストの形式が正しくありません"})
	}

	// 認証処理
	token, err := ctrl.usecase.Login(req.Username, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "認証に失敗しました")
	}

	// HttpOnly, Secure クッキーでトークンを保存
	cookie := new(http.Cookie)
	cookie.Name = "auth_token"
	cookie.Value = token
	cookie.HttpOnly = true
	cookie.Secure = true // HTTPS 環境では true にする
	cookie.Path = "/"
	c.SetCookie(cookie)

	// ユーザー情報を取得する
	user, _ := ctrl.usecase.Get(req.Username)
	userDTO := dto.UserDTO{
		ID:   user.ID,
		Name: user.Name,
		Role: user.Role,
	}

	// レスポンスを返す
	return c.JSON(http.StatusOK, map[string]any{
		"message": "ログイン成功",
		"user":    userDTO,
	})
}
