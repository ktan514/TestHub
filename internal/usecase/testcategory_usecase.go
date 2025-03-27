// ====================================
// ファイル名: testitem_usecase.go
// 作成日: 2025/03/27
// 作成者: K.Takeuchi
// 説明: このファイルの概要を記述
// ====================================

package usecase

import "testhub/internal/domain"

type TestCategoryUsecase interface {
	GetAll() ([]domain.TestCategory, error)
	Create(item domain.TestCategory) error
}

type testCategoryUsecase struct {
	repo domain.TestCategoryRepository
}

func NewTestCategoryUsecase(r domain.TestCategoryRepository) TestCategoryUsecase {
	return &testCategoryUsecase{repo: r}
}

func (u *testCategoryUsecase) GetAll() ([]domain.TestCategory, error) {
	return u.repo.FindAll()
}

func (u *testCategoryUsecase) Create(item domain.TestCategory) error {
	return u.repo.Save(item)
}
