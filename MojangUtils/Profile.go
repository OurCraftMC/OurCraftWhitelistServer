package MojangUtils

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

func GetProfileByUserName(name string) (profile ProfileStruct, err error) {
	url := "https://api.mojang.com/users/profiles/minecraft/"
	url += name
	profile = ProfileStruct{}

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return profile, err
	}

	if resp.StatusCode != 200 {
		log.Printf("GetProfileByUserName status code error: %d %s", resp.StatusCode, resp.Status)
		return profile, errors.New("not found")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	//解析json

	err = json.Unmarshal(body, &profile)
	if err != nil {
		return profile, err
	}

	return profile, nil
}
