package entities

type User71 struct {
	Name  string
	email string
}

type user struct {
	Name string
	Email string
}

type Admin struct {
	user
	Rights int
}

