package main

import (
	"dmfed/plainauth"
	"net/http"
)

func testpage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK, you're authorized now"))
}

func main() {
	wrapper := plainauth.New("yoda")
	authtestpage := wrapper.WrapHandlerFunc(testpage)
	http.HandleFunc("/normal", testpage)
	http.HandleFunc("/auth", authtestpage)
	http.ListenAndServe("127.0.0.1:10000", nil)
}
