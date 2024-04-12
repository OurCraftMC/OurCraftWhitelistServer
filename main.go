package main

import (
	"WhiteListServer/WhitelistUtils"
	"WhiteListServer/handler"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	WhitelistUtils.ReadWhitelistFromJsonFile("D://whitelist.json")

	http.HandleFunc("/api/applywhitelist", handler.ApplyWhitelist)

	log.Println("Start server")
	http.ListenAndServe(":8080", nil)
}
