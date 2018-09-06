package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)
//定义结构体
type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}
//实现接口函数
//定义为指针类型的 接收者
func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(
		resp, true)

	resp.Body.Close() // 需要关闭

	if err != nil {
		panic(err)
	}

	return string(result)
}
