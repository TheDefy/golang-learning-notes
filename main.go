package hello

import (
	"fmt"
	"strings"
	"time"
	"bytes"
	"io"
	"os"
)

func main() {

	fmt.Println("hello word!")

	//create()

	//createFunc()

	//trimTest("lei")
	bufferTest()
}

var array = [5]int{1, 2, 3, 4, 5}

var slice1 []int

func create() {
	slice5 := make([]int, 5)

	for index, value := range slice5 {
		fmt.Println("before", index, value)
	}

	fmt.Println(cap(slice5))
	slice5[0] = 1
	slice5[1] = 2
	slice5[2] = 3
	slice5[3] = 4
	slice5[4] = 5

	for index, value := range slice5 {
		fmt.Println("after", index, value)
	}
	slice := array[1:3]

	for index, value := range slice {
		fmt.Printf("%d,%d", index, value)
	}
}

type user struct {
	name  string
	email string
}

func (user *user) changeEmail(email string) {
	user.email = email
}

func (user user) notify() {
	fmt.Printf("发送邮件给%s<%s>\n", user.name, user.email)
}

func createFunc() {

	chengLei := user{"chenglei", "thedefy@163.com"}
	chengLei.notify()

	chengLei.changeEmail("TheDefy@163.com")
	chengLei.notify()
	(&chengLei).changeEmail("%TheDefy@163.com")
	chengLei.notify()

}

func trimTest(s string) string {

	trim := strings.Trim("chenglei hh lei ni hao lei lei ", s)

	fmt.Println("after string : %s", trim)

	fmt.Println(isShellSpecialVar('8'))

	now := time.Now()
	fmt.Println(now)
	add := now.Add(2000)
	fmt.Println(add)

	return trim
}

func isShellSpecialVar(c uint8) bool {

	switch c {

	case '*', '#', '$', '@', '5':
		return true

	}
	return false
}

func bufferTest()  {
	var b bytes.Buffer
	b.Write([]byte("hello"))// 将字符串写入到buffer中
	fmt.Fprintf(&b,"word")

	io.Copy(os.Stdout,&b)

}
