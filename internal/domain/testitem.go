// ====================================
// ファイル名: testitem.go
// 作成日: 2025/03/27
// 作成者: K.Takeuchi
// 説明: このファイルの概要を記述
// ====================================

package domain

type TestItem struct {
	ID              string
	Situation       string
	Purpose         string
	Operation       string
	ExpResult       string
	Topic           string
	Category        TestCategory
	ScheduledTester User
}

type TestItemRepository interface {
	FindAll() ([]TestItem, error)
	FindByID(id string) (TestItem, error)
	FindByCategory(id int) ([]TestItem, error)
	Save(item TestItem) error
	Delete(id string) error
}
