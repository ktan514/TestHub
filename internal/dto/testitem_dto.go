// ====================================
// ファイル名: testitem_dto.go
// 作成日: 2025/03/28
// 作成者: K.Takeuchi
// 説明: このファイルの概要を記述
// ====================================

package dto

type TestItemDTO struct {
	ID              int             `json:"id"`
	Subject         string          `json:"subject"`
	Perspective     string          `json:"perspective"`
	Conditions      string          `json:"conditions"`
	Steps           string          `json:"steps"`
	ExpectedResult  string          `json:"expectedResult"`
	ScheduledTester UserDTO         `json:"scheduledTester"`
	Category        TestCategoryDTO `json:"category"`
}
