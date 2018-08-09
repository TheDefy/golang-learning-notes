package hello

import "fmt"

type user60 struct {
	name  string
	email string
}

type notifier60 interface {
	notify60()
}

func (user60 *user60) notify60() {
	fmt.Println("user send email to %s<%s>",
		user60.name,
		user60.email)
}

type admin60 struct {
	user60 // 潜入类型
	level string
}

func (admin60 *admin60) notify60() {
	fmt.Println("admin send email to %s(%s)",
		admin60.name,
		admin60.email)
}

func main() {
	ad := admin60{
		user60: user60{
			name:  "ChengLei",
			email: "TheDefy@163.com",
		},
		level: "super",
	}

	sendNotification60(&ad)
	sendNotification60(&ad.user60)
}

func sendNotification60(n notifier60) {
	n.notify60()
}
