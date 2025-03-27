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

type TestItemRepository struct {
	data []domain.TestItem
	mu   sync.Mutex
}

func NewTestItemRepository() domain.TestItemRepository {
	return &TestItemRepository{
		data: []domain.TestItem{
			{ID: 1, Subject: "ログインテスト"},
			{ID: 2, Subject: "ファイルアップロードテスト"},
		},
	}
}

func (r *TestItemRepository) FindAll() ([]domain.TestItem, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.data, nil
}

func (r *TestItemRepository) FindByCategory(id int) ([]domain.TestItem, error) {
	filtered := []domain.TestItem{}
	for _, item := range r.data {
		if item.Category.ID == id {
			filtered = append(filtered, item)
		}
	}
	return filtered, nil
}

func (r *TestItemRepository) Save(item domain.TestItem) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	item.ID = len(r.data) + 1
	r.data = append(r.data, item)
	return nil
}
