package main

import (
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"u2pppw/errhandling/filelistingserver/filelisting"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

func errWrapper(
	handler appHandler) func(
	http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter,
		request *http.Request) {
		// panic
		//自己处理panic
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)

		if err != nil {
			log.Printf("Error occurred "+
				"handling request: %s",
				err.Error())

			// user error
			if userErr, ok := err.(userError); ok {
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}

			// system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
				//无权限错误
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				//不知道什么错
				code = http.StatusInternalServerError
			}
			http.Error(writer,
				http.StatusText(code), code)
		}
	}
}

//可以给用户看的error
type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/list/",
		func(writer http.ResponseWriter, request *http.Request){
			path := request.URL.Path[len("/list/"):]
			file, err := os.Open(path)
			if err != nil{
				//panic(err)
				http.Error(writer,
					err.Error(),
					http.StatusInternalServerError)
				//不能缺少return;
				return
			}
			defer file.Close()

			all,err := ioutil.ReadAll(file)
			if err != nil {
				panic(err)
			}
			writer.Write(all)
		})

	http.HandleFunc("/",
		errWrapper(filelisting.HandleFileList))
	//开启服务器
	//localhost:8888/list/fib.txt
	//http.HandleFunc("/list/",
	//	filelisting.HandleFileList)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
