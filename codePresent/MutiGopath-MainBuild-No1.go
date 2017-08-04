// +build OMIT
package main

import "fmt"

func main() {
	printHello()
}

func printHello() {
// OMIT START 選項 A：會將可執行檔建至 $HOME/go/bin 內。 // 1 OMIT END
	fmt.Println("→ 答錯了！")
	return
}
