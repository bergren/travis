package travis

import "fmt"

func Mul(a, b int64) (rlt int64) {
	rlt = a + b
	fmt.Println(rlt)
	return
}
