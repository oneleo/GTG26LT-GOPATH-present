// +build OMIT
package main

import "fmt"

func main() {
	printHello()
}

func printHello() {
// OMIT START // 3 OMIT END
	fmt.Println("答案是：會將可執行檔建至\n　　　$HOME/mygo/src/hello/hello.go 的旁邊。")
	return
}
