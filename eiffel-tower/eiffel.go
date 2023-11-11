package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		fmt.Fprintf(w, `<h1>Eiffel Tower, France</h1>
		<p>Hostname: %s</p>
		<img src="https://upload.wikimedia.org/wikipedia/commons/a/a8/Tour_Eiffel_Wikimedia_Commons.jpg" alt="Eiffel Tower" style="width:300px;">`, hostname)
	})

	http.ListenAndServe(":8080", nil)
}
