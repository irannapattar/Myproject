package Hello

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"
)

func Helloworld() {
	total := hello()
	fmt.Println("Total", total)
	l := getLenofString("hello")
	fmt.Println(l)
	sayHello()

	sayHello()
	port := getPortNum()
	fmt.Printf("TCP Port %q is available\n", port)
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
func getPortNum() string {
	low := 1000
	high := 99999
	for {
		src := rand.NewSource(time.Now().UnixNano())
		pt := low + rand.New(src).Intn(high)
		fmt.Println(pt)
		port := strconv.Itoa(pt)

		ln, err := net.Listen("tcp", ":"+port)

		if err != nil {
			fmt.Println("port running")
			continue
		}

		_ = ln.Close()

		return port
	}
}
