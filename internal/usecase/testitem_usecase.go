// ====================================
// ファイル名: testitem_usecase.go
// 作成日: 2025/03/27
// 作成者: K.Takeuchi
// 説明: このファイルの概要を記述
// ====================================

package usecase

import "testhub/internal/domain"

type TestItemUsecase interface {
	GetAll() ([]domain.TestItem, error)
	GetByCategory(id int) ([]domain.TestItem, error)
	Create(item domain.TestItem) error
}

type testItemUsecase struct {
	repo domain.TestItemRepository
}

func NewTestItemUsecase(r domain.TestItemRepository) TestItemUsecase {
	return &testItemUsecase{repo: r}
}

func (u *testItemUsecase) GetAll() ([]domain.TestItem, error) {
	return u.repo.FindAll()
}

func (u *testItemUsecase) GetByCategory(id int) ([]domain.TestItem, error) {
	return u.repo.FindByCategory(id)
}

func (u *testItemUsecase) Create(item domain.TestItem) error {
	return u.repo.Save(item)
}
