// ====================================
// ファイル名: jwt.go
// 作成日: 2025/03/31
// 作成者: K.Takeuchi
// 説明: このファイルの概要を記述
// ====================================
package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWT用の秘密鍵（環境変数などで管理）
var secretKey = []byte("your_secret_key")

// JWTトークン生成
func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // 1時間後にトークンが切れる
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// JWTトークンの検証
func ValidateJWT(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// 署名方法がHS256かどうかを確認
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
