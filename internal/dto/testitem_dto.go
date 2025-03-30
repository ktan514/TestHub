// ====================================
// ファイル名: testitem_dto.go
// 作成日: 2025/03/28
// 作成者: K.Takeuchi
// 説明: このファイルの概要を記述
// ====================================

package dto

type TestItemDTO struct {
	ID              string          `json:"id"`
	Situation       string          `json:"situation"`
	Purpose         string          `json:"purpose"`
	Operation       string          `json:"operation"`
	ExpResult       string          `json:"expResult"`
	Topic           string          `json:"topic"`
	Category        TestCategoryDTO `json:"category"`
	ScheduledTester UserDTO         `json:"scheduledTester"`
}
