package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

// 函数是一等公民

var count = 0

type handlerFunc func(http.ResponseWriter, *http.Request)

type testFunc func(int)

func addntest(n int) {
	count += n
}

func (this handlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	this(w, r)
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hello handler."))
}

func main() {
	//
	hf := handlerFunc(helloHandler)
	resp := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", bytes.NewBuffer([]byte("test")))

	hf.ServeHTTP(resp, req)

	bts, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bts))

	fmt.Println(count)
	tf := testFunc(addntest)
	tf(5)
	fmt.Println(count)
}
