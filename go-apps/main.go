package main

import (
	"fmt"
	"sync"
	"vdt-vcs-be/go-apps/mathutil"
	// "os"
	// "strconv"
)

// func sum(first, second float64) float64 {
// 	return first + second
// }

// func divide(first, second float64) (float64, error) {
// 	if second == 0 {
// 		return 0, fmt.Errorf("KHÔNG thể chia cho 0")
// 	}

// 	return first / second, nil
// }

// func main() {
// 	if len(os.Args) != 4 {
// 		fmt.Println("Cú pháp: go run . <sum|divide> <num> <num>")
// 		return
// 	}

// 	first, err := strconv.ParseFloat(os.Args[2], 64)
// 	if err != nil {
// 		fmt.Printf("Số không hợp lệ: %s\n", os.Args[2])
// 		return
// 	}

// 	second, err := strconv.ParseFloat(os.Args[3], 64)
// 	if err != nil {
// 		fmt.Printf("Số không hợp lệ: %s\n", os.Args[3])
// 		return
// 	}

// 	switch os.Args[1] {
// 	case "sum":
// 		fmt.Printf("Tổng: %g\n", sum(first, second))
// 	case "divide":
// 		result, err := divide(first, second)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		fmt.Printf("Thương: %g\n", result)
// 	default:
// 		fmt.Printf("LỆNH KHÔNG ĐÚNG: %s\n", os.Args[1])
// 	}
// }


// func worker(result chan<- string) {
// 	result <- "worker finished"
// }

// func main() {
// 	result := make(chan string)
// 	go worker(result)
// 	fmt.Println(<-result)
// }
// package main

// import (
// 	"fmt"
// 	"sync"
// )



var (
	// Mutex dùng để khóa vùng truy cập counter
	mu      sync.Mutex 
	// Biến dùng chung giữa các goroutine
	counter int        
)

func increment(wg *sync.WaitGroup) {
	// báo WaitGroup khi goroutine hoàn thành
	defer wg.Done() 

	// Khóa, không cho goroutine khác truy cập counter
	mu.Lock()       
	// Thay đổi dữ liệu dùng chung
	counter++       
	// Mở khóa
	mu.Unlock()     
}

func main() {
	// var wg sync.WaitGroup

	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go increment(&wg)
	// }

	// wg.Wait()

	// fmt.Println(counter)

	fmt.Println("Max:", mathutil.Max(10, 25))
}