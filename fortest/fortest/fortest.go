package fortest

import (
	"strconv"
	"os"
	"bufio"
	"fmt"
	"reflect"
	"runtime"
)

// 数字转二进制
func ConvertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lbs := n % 2
		result = strconv.Itoa(lbs) + result
	}
	return result
}

func PrintFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		println(scanner.Text())
	}

}

func Div(a, b int) (商, 余数 int) {
	商 = a / b
	余数 = a % b
	return
}

func Apply(op func(int, int) int, a, b int) int {
	fmt.Printf("被调用的函数名：%s  参数：%d %d \n",
		runtime.FuncForPC(reflect.ValueOf(op).Pointer()).Name(),
		a, b)
	return op(a, b)
}

func Sum(numbers ...int) int {
	var s int
	for i := range numbers {
		s += numbers[i]
	}
	return s

}
