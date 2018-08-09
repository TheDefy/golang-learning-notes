package base

import (
	"fmt"
	"strings"
	"time"
	"bytes"
	"io"
	"os"
)

var array = [5]int{1, 2, 3, 4, 5}

var slice1 []int

func Create() {
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

type User struct {
	name  string
	email string
}

func (user *User) changeEmail(email string) {
	user.email = email
}

func (user User) notify() {
	fmt.Printf("发送邮件给%s<%s>\n", user.name, user.email)
}

func CreateFunc() {

	chengLei := User{"chenglei", "thedefy@163.com"}
	chengLei.notify()

	chengLei.changeEmail("TheDefy@163.com")
	chengLei.notify()
	(&chengLei).changeEmail("%TheDefy@163.com")
	chengLei.notify()

}

func TrimTest(s string) string {

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

func BufferTest() {
	var b bytes.Buffer
	b.Write([]byte("LearningNotes ")) // 将字符串写入到buffer中
	fmt.Fprintf(&b, "world !")

	io.Copy(os.Stdout, &b)
}
