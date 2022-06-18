/**
 * @Author: realpeanut
 * @Date: 2022/6/18 07:59
 */
package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

type fooHandler struct {
}

func (f fooHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte{'h', 'e', 'l', 'l', 'o'})
	if err != nil {
		return
	}
}

func main() {

	http.Handle("/foo", fooHandler{})
	http.HandleFunc("/bar", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintf(writer, "Hello,%q", html.EscapeString(request.URL.Path))
		if err != nil {
			return
		}
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
