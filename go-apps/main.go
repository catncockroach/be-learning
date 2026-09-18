// Chương trình chính: chạy lần lượt các bài học thực hành Golang.
package main

import (
	"log"

	// "go-apps/bai1golang"
	"go-apps/bai2golangpostgresql"
)

func main() {
	// fmt.Println("########## Bài học 1 - Thực hành Golang ##########")
	// bai1golang.Run()

	if err := bai2golangpostgresql.Run(); err != nil {
		log.Println("Lỗi 2. Golang - PostgreSQL", err)
	}
}
