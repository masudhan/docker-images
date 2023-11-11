package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		fmt.Fprintf(w, `<h1>Statue of Liberty, USA</h1>
		<p>Hostname: %s</p>
		<img src="https://images.unsplash.com/photo-1503572327579-b5c6afe5c5c5?q=80&w=1942&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D" alt="Statue of Liberty" style="width:300px;">`, hostname)
	})

	http.ListenAndServe(":8080", nil)
}
