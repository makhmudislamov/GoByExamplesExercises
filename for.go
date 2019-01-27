// package main

// import "fmt"

// //for is only loop in GoLang & there r 3 types

// func main() {

// 	// 1-basic type >> single condition for
// 	i := 1
// 	for i <= 3 {
// 		fmt.Println(i)
// 		i = i + 1
// 		// prints 123, each in new line
// 	}

// 	// 2-classic for >> initial/condition/after
// 	for j := 7; j <= 9; j++ {
// 		fmt.Println(j)
// 		// prints 789, each in new line
// 	}

// 	// 3-without condition for loop >> runs until "break"
// 	// or "return" from the enclosing func
// 	for {
// 		fmt.Println("Loop")
// 		// breaks here
// 		break
// 	}
// 	for n := 0; n <= 5; n++ {
// 		if n%2 == 0 {
// 			continue
// 		}
// 		// prints 135 and exits
// 		fmt.Println(n)
// 	}

// }
