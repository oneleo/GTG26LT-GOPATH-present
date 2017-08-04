// +build OMIT
package main

import "fmt"

func main() {
	printHello()
}

func printHello() {
// OMIT START 選項 B：會將第三方套件下載到 $HOME/mygo/src 內。 // OMIT END
	fmt.Println("→ 答錯了！")
	return
}