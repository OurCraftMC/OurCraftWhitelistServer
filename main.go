package main

import (
	"WhiteListServer/Config"
	"WhiteListServer/WhitelistUtils"
	"WhiteListServer/handler"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	if len(os.Args) < 2 {
		fmt.Print("Usage: WhiteListServer [config file]\n")
		os.Exit(1)
	}
	log.Printf("Loading Config File: %s", os.Args[1])
	err := Config.LoadConfig(os.Args[1])
	if err != nil {
		log.Fatalf("Load Config File Failed: %v", err)
	}

	WhitelistUtils.ReadWhitelistFromJsonFile(Config.GetConfig().WhitelistFilePath)

	// ctrl+c 保存
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, os.Kill)
	go func() {
		<-sigs
		log.Println("Received Kill Signal, Saving files")
		WhitelistUtils.SaveWhitelistToJsonFile(Config.GetConfig().WhitelistFilePath)
		os.Exit(0)
	}()

	// handler
	http.HandleFunc("/api/applywhitelist", handler.ApplyWhitelist)
	http.HandleFunc("/", handler.MultiLanguageIndex)
	log.Println("Start server")
	http.ListenAndServe(":8080", nil)
}
