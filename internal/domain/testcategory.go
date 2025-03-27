// ====================================
// ファイル名: testcategory.go
// 作成日: 2025/03/28
// 作成者: K.Takeuchi
// 説明: このファイルの概要を記述
// ====================================

package domain

type TestCategory struct {
	ID   int
	Name string
}

type TestCategoryRepository interface {
	FindAll() ([]TestCategory, error)
	Save(item TestCategory) error
}
