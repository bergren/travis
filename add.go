package travis

import "fmt"

func Add(a, b int64) (rlt int64) {
	rlt = a + b
	fmt.Println(rlt)
	return
}
