// +build OMIT
package main

import "fmt"

func main() {
	printHello()
}

func printHello() { 
// OMIT START 選項 B：會將函數庫檔（*.a）建至 $HOME/mygo/pkg 內。 // 2 OMIT END
	fmt.Println("→ 還是答錯了！")
	return
}