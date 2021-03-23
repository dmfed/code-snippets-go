package main

import (
	"dmfed/plainauth"
	"fmt"
	"io"
	"net/http"
)

func main() {
	wrapper := plainauth.New("yoda")
	client := &http.Client{}
	r, err := http.NewRequest("GET", "http://127.0.0.1:10000/auth", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
	}
	d, _ := io.ReadAll(resp.Body)
	fmt.Println(string(d))
	r, err = http.NewRequest("GET", "http://127.0.0.1:10000/auth", nil)
	wrapper.WrapRequest(r)
	resp, err = client.Do(r)
	if err != nil {
		fmt.Println(err)
	}
	d, _ = io.ReadAll(resp.Body)
	fmt.Println(string(d))
}
