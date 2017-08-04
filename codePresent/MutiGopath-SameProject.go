// +build OMIT
package main

import "fmt"

func main() {
	printHello()
}

func printHello() { 
// OMIT START // OMIT END
	fmt.Println("→ 回去試試看吧！")
	return
}