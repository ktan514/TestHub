// ====================================
// ファイル名: testitem_repository.go
// 作成日: 2025/03/28
// 作成者: K.Takeuchi
// 説明: このファイルの概要を記述
// ====================================

package gateway

import (
	"fmt"
	"slices"
	"sync"
	"testhub/internal/domain"
)

type TestItemRepository struct {
	data []domain.TestItem
	mu   sync.Mutex
}

func NewTestItemRepository() domain.TestItemRepository {
	return &TestItemRepository{
		data: []domain.TestItem{},
	}
}

func (r *TestItemRepository) FindAll() ([]domain.TestItem, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.data, nil
}

func (r *TestItemRepository) FindByID(id string) (domain.TestItem, error) {
	filtered := domain.TestItem{}
	for _, item := range r.data {
		if item.ID == id {
			filtered = item
			break
		}
	}
	return filtered, nil
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

	// 既存の項目を検索し、インデックスを取得
	for i, existingItem := range r.data {
		if existingItem.ID == item.ID {
			r.data[i] = item // 更新
			return nil
		}
	}

	// 新規追加
	r.data = append(r.data, item)
	return nil
}

func (r *TestItemRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// スライス内の対象データを探す
	for i, item := range r.data {
		if item.ID == id {
			// 該当の要素を削除（順序を維持）
			r.data = slices.Delete(r.data, i, i+1)
			return nil
		}
	}
	return fmt.Errorf("ID %s のデータが見つかりません", id)
}
