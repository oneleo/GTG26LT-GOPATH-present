// +build OMIT
package main

import "fmt"

func main() {
	printHello()
}

func printHello() {
// OMIT START 選項 A：會將函數庫檔（*.a）建至 $HOME/go/pkg 內。 // 1 OMIT END
	fmt.Println("→ 答錯了！")
	return
}