package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start program...")
	fmt.Println(say.CallFromSay() + hello.CallFrimHello())
}
