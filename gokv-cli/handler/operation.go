package handler

import (
	"encoding/json"
	"fmt"
)

func Set(key, value string) string {
	url := buildUrl(SET)
	resp, error := Client.Get(url + "?key=" + key + "&value=" + value)
	if error != nil {
		fmt.Println(error)
		return "ERROR"
	}


	//data := []byte{}
	//__, _ := resp.Body.Read(data)

	defer resp.Body.Close()

	setResp := new(SetResp)
	json.NewDecoder(resp.Body).Decode(setResp)

	//json.Unmarshal(data, setResp)
	if setResp.Code == 200 {
		return setResp.Data
	}
	return setResp.Msg
}

func Get(key string) string {
	url := buildUrl(GET)
	resp, error := Client.Get(url + "?key=" + key)
	if error != nil {
		fmt.Println(error)
		return "ERROR"
	}

	//var data []byte
	//num, _ := resp.Body.Read(data)
	defer resp.Body.Close()

	getResp := new(GetResp)

	json.NewDecoder(resp.Body).Decode(getResp)
	//json.Unmarshal(data, getResp)
	if getResp.Code == 200 {
		return getResp.Data
	}
	return getResp.Msg
}

func Exist(key string) bool {
	return false
}

func Keys() []string {
	return []string{}
}
