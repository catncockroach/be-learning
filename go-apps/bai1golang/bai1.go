// Package bai1golang - Bài học 1: Thực hành Golang cơ bản.
// Nội dung: hàm & xử lý lỗi, goroutine + WaitGroup + Mutex, dùng module nội bộ (mathutil).
package bai1golang

import (
	"fmt"
	"sync"

	"go-apps/bai1golang/mathutil"
)

// sum trả về tổng của hai số.
func sum(first, second float64) float64 {
	return first + second
}

// divide trả về thương của hai số, báo lỗi nếu chia cho 0.
func divide(first, second float64) (float64, error) {
	if second == 0 {
		return 0, fmt.Errorf("KHÔNG thể chia cho 0")
	}
	return first / second, nil
}

var (
	mu      sync.Mutex // khóa vùng truy cập counter
	counter int        // biến dùng chung giữa các goroutine
)

// increment tăng counter một cách an toàn giữa nhiều goroutine.
func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	counter++
	mu.Unlock()
}

// Run thực thi toàn bộ demo của bài học 1.
func Run() {
	fmt.Println("=== Hàm & xử lý lỗi ===")
	fmt.Println("Tổng 10 + 25:", sum(10, 25))
	if result, err := divide(10, 0); err != nil {
		fmt.Println("Lỗi:", err)
	} else {
		fmt.Println("Thương:", result)
	}

	fmt.Println("\n=== Goroutine + WaitGroup + Mutex ===")
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("Counter sau 10 goroutine:", counter)

	fmt.Println("\n=== Module nội bộ (mathutil) ===")
	fmt.Println("Max(10, 25):", mathutil.Max(10, 25))
}
