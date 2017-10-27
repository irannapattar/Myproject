package Hello

import "fmt"

func Helloworld() {
	total := hello()
	fmt.Println("Total", total)
	fmt.Println("before statement")
	fmt.Println("end of statement")
	l := getLenofString("hello")
	fmt.Println(l)
	sayHello()
}

func hello() int {
	total := 0
	for i := 0; i < 5; i++ {
		total += i
	}
	fmt.Println("total", total)
	return total
}
func getLenofString(s string) int {
	return len(s)
}
func sayHello() {
	fmt.Println("hello")
}
