package QQGroupUtils

import (
	"WhiteListServer/Config"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func CheckIfUserInGroup(qq string) (bool, string, error) {
	url := Config.GetConfig().MiraiAddr + "get_group_member_info?group_id=" + Config.GetConfig().GroupID + "&user_id=" + qq
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("CheckIfUserInGroup: Error when get group member info,%v", err)
		return false, "", err
	}
	defer resp.Body.Close()
	info := MiraiReplyStruct{}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("CheckIfUserInGroup: Error when read group member info,%v", err)
		return false, "", err
	}
	err = json.Unmarshal(body, &info)
	if err != nil {
		log.Printf("CheckIfUserInGroup: Error when unmarshal group member info,%v", err)
		return false, "", err
	}

	if info.RetCode == 0 {
		log.Printf("CheckIfUserInGroup: User %s is in group, nickname %s , card %s", qq, info.Data.NickName, info.Data.Card)
		return true, info.Data.NickName, nil
	} else {
		log.Printf("CheckIfUserInGroup: User %s is not in group", qq)
		return false, "", nil
	}

}
