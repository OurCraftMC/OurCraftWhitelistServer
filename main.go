package main

import (
	"WhiteListServer/WhitelistUtils"
	"WhiteListServer/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	WhitelistUtils.ReadWhitelistFromJsonFile("D://whitelist.json")

	// ctrl+c 保存
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, os.Kill)
	go func() {
		<-sigs
		log.Println("Received Kill Signal, Saving files")
		WhitelistUtils.SaveWhitelistToJsonFile("D://whitelist.json")
		os.Exit(0)
	}()

	// handler
	http.HandleFunc("/api/applywhitelist", handler.ApplyWhitelist)
	http.Handle("/", http.FileServer(http.Dir("./html/")))
	log.Println("Start server")
	http.ListenAndServe(":8080", nil)
}
