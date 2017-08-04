// +build OMIT
package main

import "fmt"

func main() {
	printHello()
}

func printHello() {
// OMIT START 選項 A：會將第三方套件下載到 $HOME/go/src 內。 // OMIT END
	fmt.Println("→ 答對了！")
	return
}
