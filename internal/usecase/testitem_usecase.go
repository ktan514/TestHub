// ====================================
// ファイル名: testitem_usecase.go
// 作成日: 2025/03/27
// 作成者: K.Takeuchi
// 説明: このファイルの概要を記述
// ====================================

package usecase

import (
	"fmt"
	"testhub/internal/domain"
)

type TestItemUsecaseInterface interface {
	GetAll() ([]domain.TestItem, error)
	GetByID(id string) (domain.TestItem, error)
	GetByCategory(id int) ([]domain.TestItem, error)
	Create(item domain.TestItem) error
	Update(item domain.TestItem) error
	Delete(id []string) error
}

type TestItemUsecase struct {
	repo domain.TestItemRepository
}

func NewTestItemUsecase(r domain.TestItemRepository) TestItemUsecaseInterface {
	return &TestItemUsecase{repo: r}
}

func (u *TestItemUsecase) GetAll() ([]domain.TestItem, error) {
	return u.repo.FindAll()
}

func (u *TestItemUsecase) GetByID(id string) (domain.TestItem, error) {
	return u.repo.FindByID(id)
}

func (u *TestItemUsecase) GetByCategory(id int) ([]domain.TestItem, error) {
	return u.repo.FindByCategory(id)
}

func (u *TestItemUsecase) Create(item domain.TestItem) error {
	return u.repo.Save(item)
}

func (u *TestItemUsecase) Update(item domain.TestItem) error {
	return u.repo.Save(item)
}

func (u *TestItemUsecase) Delete(ids []string) error {
	for _, id := range ids {
		if err := u.repo.Delete(id); err != nil {
			return fmt.Errorf("ID %s の削除に失敗しました: %w", id, err)
		}
	}
	return nil
}
