package main

import "fmt"

// #include "main.h"
import "C"

func main() {
	fmt.Printf("Hello Go!\n")
	C.sayHello()
}
