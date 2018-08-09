package hello

import "fmt"

type user02 struct {
	name string

	email string
}

func (user02 *user02) notify02() {
	fmt.Printf("发送邮件给%s<%s>",
		user02.name,
		user02.email)
}

type admin02 struct {
	user02 //潜入类型

	level string
}

func (admin02 *admin02) notify02() {
	fmt.Printf("admin02发送邮件给%s(%s)",
		admin02.name,
		admin02.email)
}

type notifier01 interface {
	notify02()
}

func main() {
	ad := admin02{
		user02: user02{
			name:  "ChengLei",
			email: "TheDefy@163.com",
		},
		level: "1",
	}
	sendNotification(&ad)
	fmt.Println()
	ad.user02.notify02()
	fmt.Println()
	ad.notify02()

}

func sendNotification(n notifier01) {
	n.notify02()
}
