// ====================================
// ファイル名: main.go
// 作成日: 2025/03/27
// 作成者: K.Takeuchi
// 説明: このファイルの概要を記述
// ====================================

package main

import (
	"log"
	"testhub/infrastructure"
)

func main() {
	e := infrastructure.NewRouter()
	log.Println("Starting server on :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
