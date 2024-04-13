package handler

import (
	"WhiteListServer/WhitelistUtils"
	"encoding/json"
	"log"
	"net/http"
)

// 从一个接口获取WHitelist Json
func GetWhitelistJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	j, err := json.Marshal(WhitelistUtils.WhitelistedUser)
	if err != nil {
		log.Printf("[IP:%s]Error when marshal whitelist,%v", getRealIP(r), err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(j)
}
