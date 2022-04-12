package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("hello world")
	fmt.Println("今の時刻：" + t.String())

	port := "8080"

	http.Handle("/", http.FileServer(http.Dir("root/"))) // ①
	log.Printf("Server listening on http://localhost:%s/", port)
	log.Print(http.ListenAndServe(":"+port, nil))

}
