package awesomeProject1

import "fmt"

func main() {
	go func() {
		fmt.Println("下山的路又堵起了")
	}()
}
