package mock

import "fmt"

type Retriever struct {
	Contents string
}

//实现了.String()的接口
func (r *Retriever) String() string {
	return fmt.Sprintf(
		"Retriever: {Contents=%s}", r.Contents)
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

//实现了Retriever;
func (r *Retriever) Get(url string) string {
	return r.Contents
}
