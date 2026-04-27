package main

import "fmt"

type MyError struct {
	Msg string
}

func (e *MyError) Error() string {
	return e.Msg
}

func test() error {
	var err *MyError = nil
	return err
}

func test2() error {
	return nil
}

func main() {
	err1 := test()
	err2 := test2()

	fmt.Println(err1 == nil) // ???
	fmt.Println(err2 == nil) // ???
	fmt.Println(err1)        // ???

	if err1 != nil {
		fmt.Println("Error1 not nil") // Выведется?
	}

	var err3 error
	var err4 *MyError = nil
	err3 = err4

	fmt.Println(err3 == nil) // ???
}
