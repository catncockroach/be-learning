// Package bai2golangpostgresql - Bài học 2: Thực hành Golang - PostgreSQL.
// Nội dung: kết nối PostgreSQL và thao tác CRUD bằng GORM.
package bai2golangpostgresql

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User map với bảng "users" thông qua GORM.
// GORM tự động dùng tên struct (số nhiều, snake_case) làm tên bảng nếu không chỉ định khác.
type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:100;not null"`
	Email string `gorm:"size:100;uniqueIndex"`
}

// getEnv trả về giá trị biến môi trường, hoặc fallback nếu chưa được thiết lập.
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// Run thực thi toàn bộ demo của bài học 2.
func Run() error {
	// Nạp file .env nếu có (bỏ qua lỗi để vẫn chạy được khi biến môi trường
	// đã được thiết lập sẵn, ví dụ trên môi trường production).
	_ = godotenv.Load()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "hai"),
		getEnv("DB_PASSWORD", "hai"),
		getEnv("DB_NAME", "hai"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_SSLMODE", "disable"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("không thể kết nối database: %w", err)
	}
	fmt.Println("Kết nối PostgreSQL thành công!")

	// AutoMigrate tạo/cập nhật bảng "users" theo struct User.
	if err := db.AutoMigrate(&User{}); err != nil {
		return fmt.Errorf("migrate lỗi: %w", err)
	}

	// Create
	user := User{Name: "Nguyễn Văn A", Email: "a@example.com"}
	db.Create(&user)
	fmt.Println("Create:", user)

	// Read
	var found User
	db.First(&found, "email = ?", "a@example.com")
	fmt.Println("Read:", found)

	// Update
	db.Model(&found).Update("Name", "Nguyễn Văn A (updated)")
	fmt.Println("Update:", found)

	// Danh sách tất cả user
	var users []User
	db.Find(&users)
	fmt.Println("Find all:", users)

	// Delete
	db.Delete(&found)

	var count int64
	db.Model(&User{}).Count(&count)
	fmt.Println("Số user còn lại sau khi xóa:", count)

	return nil
}
