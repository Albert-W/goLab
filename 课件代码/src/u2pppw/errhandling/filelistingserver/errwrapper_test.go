package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(_ http.ResponseWriter,
	_ *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(_ http.ResponseWriter,
	_ *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(_ http.ResponseWriter,
	_ *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(_ http.ResponseWriter,
	_ *http.Request) error {
	return os.ErrPermission
}

func errUnknown(_ http.ResponseWriter,
	_ *http.Request) error {
	return errors.New("unknown error")
}

func noError(writer http.ResponseWriter,
	_ *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}
//共用测试数据
var tests = []struct {
	// h 是上面的函数，输入
	h       appHandler
	//期望输出
	code    int
	message string
}{
	//记
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.imooc.com", nil)
		f(response, request)

		verifyResponse(response.Result(),
			tt.code, tt.message, t)
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		//起server
		server := httptest.NewServer(
			http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)

		verifyResponse(
			resp, tt.code, tt.message, t)
	}
}
//共同的验证
func verifyResponse(resp *http.Response,
	expectedCode int, expectedMsg string,
	t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	//去掉换行
	body := strings.Trim(string(b), "\n")
	//测试的判断
	if resp.StatusCode != expectedCode ||
		body != expectedMsg {
		t.Errorf("expect (%d, %s); "+
			"got (%d, %s)",
			expectedCode, expectedMsg,
			resp.StatusCode, body)
	}
}
