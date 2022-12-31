package awesomeProject1

import (
	"encoding/json"
	"fmt"
)

type People struct {
	name string
}

func main() {
	js := `{
"name":"坤坤"
}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err = nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(p)
}
