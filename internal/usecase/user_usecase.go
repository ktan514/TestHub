// ====================================
// ファイル名: testitem_usecase.go
// 作成日: 2025/03/27
// 作成者: K.Takeuchi
// 説明: このファイルの概要を記述
// ====================================

package usecase

import (
	"errors"
	"testhub/infrastructure/auth/jwt"
	"testhub/internal/domain"
)

type UserUsecaseInterface interface {
	Login(name string, password string) (string, error)
	Get(name string) (domain.User, error)
	GetAll() ([]domain.User, error)
	Create(item domain.User) error
}

type UserUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(r domain.UserRepository) UserUsecaseInterface {
	return &UserUsecase{repo: r}
}

func (u *UserUsecase) Login(name string, password string) (string, error) {
	user, err := u.repo.FindByUserName(name)
	if err != nil {
		return "", err
	}

	// パスワード検証
	if user.Password != password {
		return "", errors.New("無効なユーザー名はまたはパスワード")
	}

	// JWTトークンを生成
	token, err := jwt.GenerateJWT(name)
	if err != nil {
		return "", err
	}

	return token, err
}

func (u *UserUsecase) Get(name string) (domain.User, error) {
	return u.repo.FindByUserName(name)
}

func (u *UserUsecase) GetAll() ([]domain.User, error) {
	return u.repo.FindAll()
}

func (u *UserUsecase) Create(item domain.User) error {
	return u.repo.Save(item)
}
