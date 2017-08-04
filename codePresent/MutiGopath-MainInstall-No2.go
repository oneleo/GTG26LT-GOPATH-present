// +build OMIT
package main

import "fmt"

func main() {
	printHello()
}

func printHello() {
// OMIT START // 1 OMIT END
	fmt.Println("→ 沒有其他選項了！")
	return
}
