// ====================================
// ファイル名: user_repository.go
// 作成日: 2025/03/29
// 作成者: K.Takeuchi
// 説明: このファイルの概要を記述
// ====================================

package gateway

import (
	"errors"
	"sync"
	"testhub/internal/domain"
)

type UserRepository struct {
	data []domain.User
	mu   sync.Mutex
}

func NewUserRepository() domain.UserRepository {
	return &UserRepository{
		data: []domain.User{
			{ID: 1, Name: "Tester1", Password: "Tester1", Role: 0},
			{ID: 2, Name: "Tester2", Password: "Tester2", Role: 0},
			{ID: 3, Name: "Tester3", Password: "Tester3", Role: 0},
			{ID: 4, Name: "Tester4", Password: "Tester4", Role: 0},
			{ID: 5, Name: "Tester5", Password: "Tester5", Role: 0},
		},
	}
}

func (r *UserRepository) FindByUserName(name string) (domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, user := range r.data {
		if user.Name == name {
			return user, nil
		}
	}

	return domain.User{}, errors.New("ユーザーが見つかりませんでした。")
}

func (r *UserRepository) FindAll() ([]domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.data, nil
}

func (r *UserRepository) Save(item domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	item.ID = len(r.data) + 1
	r.data = append(r.data, item)
	return nil
}
