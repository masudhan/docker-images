package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		fmt.Fprintf(w, `<h1>Taj Mahal, India</h1>
		<p>Hostname: %s</p>
		<img src="https://images.unsplash.com/photo-1545562083-a600704fa486?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D" alt="Eiffel Tower" style="width:500px;">`, hostname)
	})

	http.ListenAndServe(":8080", nil)
}
