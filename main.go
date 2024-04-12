package main

import (
	"WhiteListServer/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/applywhitelist", handler.ApplyWhitelist)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Start server")
	http.ListenAndServe(":8080", nil)
}
