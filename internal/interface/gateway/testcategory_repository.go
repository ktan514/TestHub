// ====================================
// ファイル名: testitem_repository.go
// 作成日: 2025/03/28
// 作成者: K.Takeuchi
// 説明: このファイルの概要を記述
// ====================================

package gateway

import (
	"sync"
	"testhub/internal/domain"
)

type TestCategoryRepository struct {
	data []domain.TestCategory
	mu   sync.Mutex
}

func NewTestCategoryRepository() domain.TestCategoryRepository {
	return &TestCategoryRepository{
		data: []domain.TestCategory{
			{ID: 1, Name: "New Items"},
			{ID: 2, Name: "Normal Items"},
			{ID: 3, Name: "Function A"},
			{ID: 4, Name: "Function B"},
			{ID: 5, Name: "GUI"},
		},
	}
}

func (r *TestCategoryRepository) FindAll() ([]domain.TestCategory, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.data, nil
}

func (r *TestCategoryRepository) Save(item domain.TestCategory) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	item.ID = len(r.data) + 1
	r.data = append(r.data, item)
	return nil
}
