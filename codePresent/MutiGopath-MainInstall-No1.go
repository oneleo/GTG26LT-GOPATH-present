// +build OMIT
package main

import "fmt"

func main() {
	printHello()
}

func printHello() {
// OMIT START 選項 B：會將可執行檔建至 $HOME/mygo/bin 內。 // 2 OMIT END
	fmt.Println("→ 答錯了！")
	return
}