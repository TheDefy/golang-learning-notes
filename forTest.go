package hello

import (
	"strconv"
	"os"
	"bufio"
	"fmt"
	"reflect"
	"runtime"
	"math"
)

// 数字转二进制
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lbs := n % 2
		result = strconv.Itoa(lbs) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		println(scanner.Text())
	}

}
func main() {

	fmt.Println(
		convertToBin(3),
		convertToBin(5),
	)

	printFile("src/.gitignore")

	商, 余数 := div(13, 3)

	fmt.Println(商, 余数)

	i := apply(func(i int, i2 int) int {
		return int(math.Max(float64(i), float64(i2)))
	}, 3, 4)

	fmt.Printf("func apply rusult: %d \n", i)

	fmt.Println("sum result:", sum(1, 2, 3))
}

func div(a, b int) (商, 余数 int) {
	商 = a / b
	余数 = a % b
	return
}

func apply(op func(int, int) int, a, b int) int {
	fmt.Printf("被调用的函数名：%s  参数：%d %d \n",
		runtime.FuncForPC(reflect.ValueOf(op).Pointer()).Name(),
		a, b)
	return op(a, b)
}

func sum(numbers ...int) int {
	var s int
	for i := range numbers {
		s += numbers[i]
	}
	return s

}
