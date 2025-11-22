package main

import (
	"fmt"
	"runtime"
	"sync"
)

// import "fmt"
// import (
// 	"educationproject/src/internal/user/model"
// 	"fmt"
// 	"time"
// )

// var errorNotFoundFile = errors.New("I cannot find file")
func fanIn(ch1, ch2 <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for {
			select {
			case v, ok := <-ch1:
				if !ok {
					ch1 = nil // disable this case
					continue
				}
				out <- v

			case v, ok := <-ch2:
				if !ok {
					ch2 = nil // disable this case
					continue
				}
				out <- v
			}

			// stop when both are nil → dead channel
			if ch1 == nil && ch2 == nil {
				close(out)
				return
			}
		}
	}()

	return out
}

func main() {
	//------------- your practice
	// ch1 := make(chan int)
	// ch2 := make(chan int)

	// // sender 1
	// go func() {
	// 	for i := 0; i < 5; i++ {
	// 		ch1 <- i
	// 		time.Sleep(200 * time.Millisecond)
	// 	}
	// 	close(ch1)
	// }()

	// // sender 2
	// go func() {
	// 	for i := 100; i < 105; i++ {
	// 		ch2 <- i
	// 		time.Sleep(350 * time.Millisecond)
	// 	}
	// 	close(ch2)
	// }()

	// // merged output
	// out := fanIn(ch1, ch2)

	// // consumer
	// for v := range out {
	// 	fmt.Println("Received:", v)
	// }

	//------------- your practice

	// file, err := os.Create("tracing.out")

	// if err != nil {
	// 	println("it has some error")
	// }

	// defer file.Close()

	// trace.Start(file)
	// var wg sync.WaitGroup
	// ch := make(chan int, 1)
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		defer wg.Done()
	// 		println("Goroutine", i, "sending")
	// 		ch <- i
	// 		println("Goroutine", i, "sent")
	// 	}(i)
	// }

	// go func() {
	// 	wg.Wait()
	// }()

	// for i := 0; i < 10; i++ {
	// 	val := <-ch
	// 	println("Received", val)
	// 	time.Sleep(500 * time.Millisecond)
	// }

	// defer trace.Stop()

	// serTest := Service()

	// inner := errors.Unwrap(serTest)
	// fmt.Println(inner)
	// if errors.Is(serTest, errorNotFoundFile) {
	// 	fmt.Println("Not found file")
	// }
	// file, err := os.Open("../../data/testdata.csv")

	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("recover my error %w", err)
	// 	}

	// }()
	// if err != nil {
	// 	panic(err)
	// }

	// defer file.Close()

	// ctx := context.Background()

	// fileReader := repository.NewOperatorRepository(file)
	// value, err := fileReader.Get(ctx)

	// if err != nil {
	// 	panic(err)
	// }
	// for _, p := range value {
	// 	fmt.Printf("%+v\n", p)
	// }

	// user := model.NewUser("Morteza", "Rezaei", time.Now(), true)

	// fmt.Println(user.FirstName)
	// var user model.User

	// user.FirstName = "Test"

	// user.Modified(&user)

	// fmt.Println(user.FirstName)
	// value := 10
	// Test(&value)
	// println(value)

	// if returnTest := Test("sunday"); returnTest != "" {
	// 	fmt.Println("test is good")
	// }

	// fmt.Println("test is worst")
	var wg sync.WaitGroup
	wg.Add(2)
	defer wg.Done()

	go traceGoroutine(1)
	go traceGoroutine(2)

	wg.Wait()
}

func traceGoroutine(val int) {
	runtime.LockOSThread() // Bind goroutine to current OS thread
	defer runtime.UnlockOSThread()

	fmt.Printf("Running on OS Thread:%d\n", val) // just for illustration
}

// func Service() error {
// 	return fmt.Errorf("this is service error %w", Service2())
// }
// func Service2() error {
// 	return errorNotFoundFile
// }

// -------- why I could not run after panic

// defer func() {
// 	if r := recover(); r != nil {
// 		println("test")
// 	}
// }()

// TestPanic()

// println("after panic")

// func TestPanic() {
// 	println("before panic")
// 	panic("test")
// }

//-------- why I could not run after panic

// func Test(value *int) {
// 	println("test")
// 	*value = 20
// }

// func Test(input string) string {
// 	if input == "sunday" {
// 		fmt.Println("test is success")
// 		return "my method finished in if"
// 	}
// 	return ""

// }
