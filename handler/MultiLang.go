package handler

import (
	"log"
	"net/http"
	"strings"
)

func MultiLanguageIndex(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/" {
		http.ServeFile(w, r, "./html"+r.RequestURI)
		return
	}
	lang := r.Header.Get("Accept-Language")
	switch {
	case strings.Contains(lang, "zh"):
		log.Printf("[IP:%s]User Accept Lang: %s, Giving Chinese Page", getRealIP(r), lang)
		http.ServeFile(w, r, "./html/index.html")
	case strings.Contains(lang, "en"):
		fallthrough
	default:
		log.Printf("[IP:%s]User Accept Lang: %s, Giving English Page", getRealIP(r), lang)
		http.ServeFile(w, r, "./html/en/index.html")
	}

}
