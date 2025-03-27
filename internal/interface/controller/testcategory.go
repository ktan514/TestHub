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

type TestCategoryController struct {
	usecase usecase.TestCategoryUsecase
}

func NewTestCategoryController(u usecase.TestCategoryUsecase) *TestCategoryController {
	return &TestCategoryController{usecase: u}
}

func (ctrl *TestCategoryController) GetTestCategories(c echo.Context) error {
	items, err := ctrl.usecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// DomainからDTO(Data Transfer Object)に変換する
	testCategoryDTOs := make([]dto.TestCategoryDTO, len(items))
	for i, item := range items {
		testCategoryDTOs[i] = dto.TestCategoryDTO{
			ID:   item.ID,
			Name: item.Name,
		}
	}

	// JSONに変換して返却する
	return c.JSON(http.StatusOK, testCategoryDTOs)
}
