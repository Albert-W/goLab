package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

// 自己定义用户error类型
type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}
// 用错就返中error
func HandleFileList(writer http.ResponseWriter,
	request *http.Request) error {
	fmt.Println()
	if strings.Index(
		request.URL.Path, prefix) != 0 {
		return userError(
			fmt.Sprintf("path %s must start "+
				"with %s",
				request.URL.Path, prefix))
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}
