package main

import (
	"errors"
	"fmt"
)

// import "fmt"
// import (
// 	"educationproject/src/internal/user/model"
// 	"fmt"
// 	"time"
// )

var errorNotFoundFile = errors.New("I cannot find file")

func main() {

	serTest := Service()

	inner := errors.Unwrap(serTest)
	fmt.Println(inner)
	if errors.Is(serTest, errorNotFoundFile) {
		fmt.Println("Not found file")
	}
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
}

func Service() error {
	return fmt.Errorf("this is service error %w", Service2())
}
func Service2() error {
	return errorNotFoundFile
}

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
