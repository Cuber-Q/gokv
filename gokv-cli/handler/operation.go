package handler

import (
	"encoding/json"
	"fmt"
)

func Set(key, value string) string {
	url := buildUrl(SET)
	resp, e := Client.Get(url + "?key=" + key + "&value=" + value)
	if e != nil {
		fmt.Println(e)
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
	resp, e := Client.Get(url + "?key=" + key)
	if e != nil {
		fmt.Println(e)
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

/* cluster cmd */
// for a leader node, to add a new node as its follower
func AddNode(address string) string {
	url := buildUrl(ADD_NODE)
	resp, e := Client.Get(url + "?node=" + address)
	if e != nil {
		fmt.Println(e)
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