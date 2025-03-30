// ====================================
// ファイル名: testitem_usecase.go
// 作成日: 2025/03/27
// 作成者: K.Takeuchi
// 説明: このファイルの概要を記述
// ====================================

package usecase

import "testhub/internal/domain"

type TestCategoryUsecaseInterface interface {
	GetAll() ([]domain.TestCategory, error)
	Create(item domain.TestCategory) error
}

type TestCategoryUsecase struct {
	repo domain.TestCategoryRepository
}

func NewTestCategoryUsecase(r domain.TestCategoryRepository) TestCategoryUsecaseInterface {
	return &TestCategoryUsecase{repo: r}
}

func (u *TestCategoryUsecase) GetAll() ([]domain.TestCategory, error) {
	return u.repo.FindAll()
}

func (u *TestCategoryUsecase) Create(item domain.TestCategory) error {
	return u.repo.Save(item)
}
