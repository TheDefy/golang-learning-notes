package mock

import "fmt"

type Retriever struct {
	Contents string
}

func (r *Retriever) Post(url string) string {
	r.Contents = url
	fmt.Println("post contents:",r.Contents)
	return "ok"
}

func (r *Retriever) Get(url string) string {
	s := r.Contents
	fmt.Println("get contents: ",s)
	return s
}
