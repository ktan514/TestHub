// ====================================
// ファイル名: testitem.go
// 作成日: 2025/03/27
// 作成者: K.Takeuchi
// 説明: このファイルの概要を記述
// ====================================

package domain

type TestItem struct {
	ID              int
	Subject         string
	Perspective     string
	Conditions      string
	Steps           string
	ExpectedResult  string
	ScheduledTester string
	Category        TestCategory
}

type TestItemRepository interface {
	FindAll() ([]TestItem, error)
	FindByCategory(id int) ([]TestItem, error)
	Save(item TestItem) error
}
