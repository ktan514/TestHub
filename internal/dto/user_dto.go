// ====================================
// ファイル名: user_dto.go
// 作成日: 2025/03/28
// 作成者: K.Takeuchi
// 説明: このファイルの概要を記述
// ====================================

package dto

type UserDTO struct {
	ID   int    `json:"id"`   // ID
	Name string `json:"name"` // 名前
	Role int    `json:"role"` // 役割り
}
