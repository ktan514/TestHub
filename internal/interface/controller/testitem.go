// ====================================
// ファイル名: testitem.go
// 作成日: 2025/03/27
// 作成者: K.Takeuchi
// 説明: このファイルの概要を記述
// ====================================

package controller

import (
	"net/http"
	"strconv"
	"testhub/internal/domain"
	"testhub/internal/dto"
	"testhub/internal/usecase"

	"github.com/labstack/echo/v4"
)

type TestItemController struct {
	usecase usecase.TestItemUsecase
}

func NewTestItemController(u usecase.TestItemUsecase) *TestItemController {
	return &TestItemController{usecase: u}
}

func (ctrl *TestItemController) GetTestItems(c echo.Context) error {
	categoryId, _ := strconv.Atoi(c.QueryParam("category"))

	var items []domain.TestItem
	var err error

	items, err = ctrl.usecase.GetByCategory(categoryId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// DomainからDTO(Data Transfer Object)に変換する
	testItemDTOs := make([]dto.TestItemDTO, len(items))
	for i, item := range items {
		testItemDTOs[i] = dto.TestItemDTO{
			ID:             item.ID,
			Subject:        item.Subject,
			Perspective:    item.Perspective,
			Conditions:     item.Conditions,
			Steps:          item.Steps,
			ExpectedResult: item.ExpectedResult,
			ScheduledTester: dto.UserDTO{
				ID:   item.ScheduledTester.ID,
				Name: item.ScheduledTester.Name,
			},
			Category: dto.TestCategoryDTO{
				ID:   item.Category.ID,
				Name: item.Category.Name,
			},
		}
	}

	// JSONに変換して返却する
	return c.JSON(http.StatusOK, testItemDTOs)
}

func (ctrl *TestItemController) AddTestItem(c echo.Context) error {
	var testItemDTO dto.TestItemDTO
	if err := c.Bind(&testItemDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// DTOからDomainに変換する
	item := domain.TestItem{
		ID:             testItemDTO.ID,
		Subject:        testItemDTO.Subject,
		Perspective:    testItemDTO.Perspective,
		Conditions:     testItemDTO.Conditions,
		Steps:          testItemDTO.Steps,
		ExpectedResult: testItemDTO.ExpectedResult,
		ScheduledTester: domain.User{
			ID:   testItemDTO.ScheduledTester.ID,
			Name: testItemDTO.ScheduledTester.Name,
		},
		Category: domain.TestCategory{
			ID:   testItemDTO.Category.ID,
			Name: testItemDTO.Category.Name,
		},
	}

	// デバッグログ
	c.Logger().Infof("Received item [ID=%d, Subject='%s', Category='%s']", item.ID, item.Subject, item.Category.Name)

	err := ctrl.usecase.Create(item)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, item)
}
