package main

// import "fmt"
import (
	"educationproject/src/internal/user/model"
	"fmt"
)

func main() {
	user := model.NewUser("Morteza", "Rezaei")
	user.FirstName = "Test"

	user.Modified(user)

	fmt.Println(user.FirstName)
	// value := 10
	// Test(&value)
	// println(value)

	// if returnTest := Test("sunday"); returnTest != "" {
	// 	fmt.Println("test is good")
	// }

	// fmt.Println("test is worst")
}

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
