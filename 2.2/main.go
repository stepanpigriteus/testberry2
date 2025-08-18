package main

import "fmt"

func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return /// в функции именованный возврат, значит дефер ссначала выполнится и только потом х будет возвращена = 2
}

func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x /// в функции неименованный возврат - сначала скопированное значение x=1 будет возвращено, а затем defer выполнится
}

func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
