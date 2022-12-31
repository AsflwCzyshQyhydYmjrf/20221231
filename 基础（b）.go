package awesomeProject1

import "fmt"

func main() {
	a := 2
	b := 0
	defer func() {
		if r := recover(); r = nil {
			fmt.Println(r)
		}
	}()
	fmt.Println("2/0=", a/b)
	fmt.Println("3/2 =", 3/2)
}
