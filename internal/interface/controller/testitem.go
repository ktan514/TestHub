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
	usecase usecase.TestItemUsecaseInterface
}

func NewTestItemController(u usecase.TestItemUsecaseInterface) *TestItemController {
	return &TestItemController{usecase: u}
}

func (ctrl *TestItemController) GetTestItems(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.QueryParam("categoryId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid categoryId"})
	}

	var items []domain.TestItem

	// カテゴリーIDを条件に試験項目リストを取得する
	items, err = ctrl.usecase.GetByCategory(categoryID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// DomainからDTO(Data Transfer Object)に変換する
	testItemDTOs := make([]dto.TestItemDTO, len(items))
	for i, item := range items {
		testItemDTOs[i] = dto.TestItemDTO{
			ID:        item.ID,
			Situation: item.Situation,
			Purpose:   item.Purpose,
			Operation: item.Operation,
			ExpResult: item.ExpResult,
			Topic:     item.Topic,
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

func (ctrl *TestItemController) DelTestItems(c echo.Context) error {
	// クエリパラメータ "ids" を取得（複数指定可能）
	ids := c.QueryParams()["ids[]"]
	if len(ids) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "idsクエリパラメータが指定されていません"})
	}

	// ユースケース層でデータ削除処理を実行
	if err := ctrl.usecase.Delete(ids); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// 削除完了メッセージを返却
	return c.JSON(http.StatusOK, map[string]string{"message": "削除が完了しました"})
}

func (ctrl *TestItemController) AddTestItem(c echo.Context) error {
	var testItemDTO dto.TestItemDTO
	if err := c.Bind(&testItemDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// DTOからDomainに変換する
	item := domain.TestItem{
		ID:        testItemDTO.ID,
		Situation: testItemDTO.Situation,
		Purpose:   testItemDTO.Purpose,
		Operation: testItemDTO.Operation,
		ExpResult: testItemDTO.ExpResult,
		Topic:     testItemDTO.Topic,
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
	c.Logger().Infof("Received item [ID=%d, Test Category='%s']", item.ID, item.Category.Name)

	err := ctrl.usecase.Create(item)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, item)
}

func (ctrl *TestItemController) ModifyTestItem(c echo.Context) error {
	var testItemDTO dto.TestItemDTO
	if err := c.Bind(&testItemDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// DTOからDomainに変換する
	item := domain.TestItem{
		ID:        testItemDTO.ID,
		Situation: testItemDTO.Situation,
		Purpose:   testItemDTO.Purpose,
		Operation: testItemDTO.Operation,
		ExpResult: testItemDTO.ExpResult,
		Topic:     testItemDTO.Topic,
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
	c.Logger().Infof("Received item [ID=%d, Test Category='%s']", item.ID, item.Category.Name)

	err := ctrl.usecase.Create(item)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, item)
}

func (ctrl *TestItemController) GetTestItem(c echo.Context) error {
	id := c.QueryParam("id")

	var item domain.TestItem
	var err error

	item, err = ctrl.usecase.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// DomainからDTO(Data Transfer Object)に変換する
	testItemDTO := dto.TestItemDTO{
		ID:        item.ID,
		Situation: item.Situation,
		Purpose:   item.Purpose,
		Operation: item.Operation,
		ExpResult: item.ExpResult,
		Topic:     item.Topic,
		ScheduledTester: dto.UserDTO{
			ID:   item.ScheduledTester.ID,
			Name: item.ScheduledTester.Name,
		},
		Category: dto.TestCategoryDTO{
			ID:   item.Category.ID,
			Name: item.Category.Name,
		},
	}

	// JSONに変換して返却する
	return c.JSON(http.StatusOK, testItemDTO)
}
