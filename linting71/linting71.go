package main

import (
	"entities"
	"fmt"
)

func main() {
	/*u := entities.User{
		Name:  "chenglei",
		email: "thedefy@163.com",
	}

	fmt.Printf("User: %v\n", u)*/

	ad := entities.Admin{
		Rights: 9,
	}

	ad.Name = "ChengLei"
	ad.Email = "TheDefy@163.com"

	fmt.Printf("Admin : %v \n", ad)
}
