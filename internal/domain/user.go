// ====================================
// ファイル名: user.go
// 作成日: 2025/03/28
// 作成者: K.Takeuchi
// 説明: このファイルの概要を記述
// ====================================

package domain

type User struct {
	ID            int    // ID
	Name          string // 名前
	Role          int    // 役割り
	Administrator int    // 管理者権限フラグ
}
