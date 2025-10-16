package main

import "fmt"

func main() {
	hello := func(date string) {
		fmt.Printf("Hello, %s!\n", date)
	}

	Sayhello("3/31", hello)

}

func Sayhello(date string, hello func(date string)) {
	hello(date)
}
