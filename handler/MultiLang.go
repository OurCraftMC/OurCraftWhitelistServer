package handler

import (
	"WhiteListServer/Config"
	"log"
	"net/http"
	"strings"
)

func MultiLanguageIndex(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/" {
		http.ServeFile(w, r, Config.GetConfig().HTMLPath+r.RequestURI[1:])
		return
	}
	lang := r.Header.Get("Accept-Language")
	switch {
	case strings.Contains(lang, "zh"):
		log.Printf("[IP:%s]User Accept Lang: %s, Giving Chinese Page", getRealIP(r), lang)
		http.ServeFile(w, r, Config.GetConfig().HTMLPath+"index.html")
	case strings.Contains(lang, "en"):
		fallthrough
	default:
		log.Printf("[IP:%s]User Accept Lang: %s, Giving English Page", getRealIP(r), lang)
		http.ServeFile(w, r, Config.GetConfig().HTMLPath+"index_en.html")
	}

}
