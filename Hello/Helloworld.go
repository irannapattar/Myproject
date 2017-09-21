package Hello

import "fmt"

func Helloworld() {
	total := hello()
	fmt.Println("Total", total)
}

func hello() int {
	total := 0
	for i := 0; i < 5; i++ {
		total += i
	}
	return total
}
