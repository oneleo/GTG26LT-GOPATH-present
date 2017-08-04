// +build OMIT
package main

import "fmt"

func main() {
	printHello()
}

func printHello() {
// OMIT START // 3 OMIT END
	fmt.Println("答案是：會將可執行檔建至您執行 go build 指令的目錄中。\n例如：您在 /tmp/ 下執行 go build \"hello/\"，\n可執行檔將會出現在 /tmp/ 內。")
	return
}
