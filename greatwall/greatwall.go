package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		fmt.Fprintf(w, `<h1>Great Wall of China</h1>
		<p>Hostname: %s</p>
		<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/2/23/The_Great_Wall_of_China_at_Jinshanling-edit.jpg/1280px-The_Great_Wall_of_China_at_Jinshanling-edit.jpg" alt="Great Wall of China" style="width:500px;">`, hostname)
	})

	http.ListenAndServe(":8080", nil)
}
