package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	diskFS := os.DirFS("./static")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServerFS(diskFS)))

	http.HandleFunc("/{$}", func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile("templates/index.html")
		if err != nil {
			log.Fatal(err)
		}

		w.Write(content)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("here")
		log.Fatal(err)
	}
}
