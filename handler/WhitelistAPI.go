package handler

import (
	"WhiteListServer/MojangUtils"
	"WhiteListServer/WhitelistUtils"
	"log"
	"net/http"
)

func getRealIP(r *http.Request) string {
	if r.Header.Get("X-Real-Ip") != "" {
		return r.Header.Get("X-Real-Ip")
	}
	return r.RemoteAddr
}

func ApplyWhitelist(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	contactMethod := r.URL.Query().Get("contactmethod")
	contactID := r.URL.Query().Get("contactid")
	ip := getRealIP(r)

	//确认不在白名单钟
	switch {
	case contactMethod == "":
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{'error':'请填写联系方式','success':false}"))
		log.Printf("[IP:%s]User %s Applied,but contactMethod is empty", ip, name)
		return
	case contactID == "":
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{'error':'请填写联系方式','success':false}"))
		log.Printf("[IP:%s]User %s Applied,but contactID is empty", ip, name)
		return
	case name == "":
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{'error':'请填写名字','success':false}"))
		log.Printf("[IP:%s]User %s Applied,but name is empty", ip)
		return
	}
	log.Printf("[IP:%s]User %s Applied,name:'%s',contactMethod:'%s',contactID:'%s'", ip, name, name, contactMethod, contactID)

	//先确认是否在白名单中
	if WhitelistUtils.CheckIfInWhitelist(name) {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("{'error':'你已经在白名单申请列表中了哦，请不要重复申请','success':false}"))
		log.Printf("[IP:%s]User %s Applied,but already in the apply list", ip, name)
		return
	}

	//获取用户相关信息
	profile, err := MojangUtils.GetProfileByUserName(name)
	if err != nil {
		switch {
		case err.Error() == "not found":
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("{'error':'找不到这个用户','success':false}"))
			log.Printf("[IP:%s]User %s Applied,but user not found(via MojangUtils)", ip, name)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{'error':'Mojang服务器错误','success':false}"))
			log.Printf("[IP:%s]User %s Applied,but MojangUtils error:%s", ip, name, err)
			return
		}
	}
	log.Printf("[IP:%s]User %s ,UUID:%s", ip, name, profile.UUID)
	WhitelistUtils.AddToWhitelist(name, profile.UUID, contactMethod, contactID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{'success':true}"))
	return

}
