package hello

import (
	"reflect"
	"runtime"
	"fmt"
	"math"
)

func applyFunc(op func(int, int) int, a, b int) int {

	pointer := reflect.ValueOf(op).Pointer()
	pc := runtime.FuncForPC(pointer)
	funName := pc.Name()

	fmt.Printf("被调用的函数名： %s  参数： %d,%d\n", funName, a, b)

	return op(a, b)
}

func maxNumber(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func main() {
	result := applyFunc(maxNumber, 3, 4)
	fmt.Println("applyFunc result:", result)
}
