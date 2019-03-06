package youdaoclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const ERROR_CODE = -1

type YouDao struct {
	Translation []string `json:"translation"`
	basic       Basic    `json:"basic"`
	Query       string   `json:"query"`
	ErrorCode   int      `json:"errorCode"`
	Web         []*Web   `json:"web"`
}

type Basic struct {
	PhoneticUS string   `json:"us-phonetic"`
	Phonetic   string   `json:"phonetic"`
	PhoneticUK string   `json:"uk-phonetic"`
	Explains   []string `json:"explains"`
}

type Web struct {
	Key   string   `json:"key"`
	Value []string `json:"value"`
}

func Translation(word string) YouDao {
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := client.Get("http://fanyi.youdao.com/openapi.do?keyfrom=SillyHatYouDao&key=987724779&type=data&doctype=json&version=1.1&q=" + word)
	if err != nil {
		fmt.Println(err)
		return *&YouDao{ErrorCode: ERROR_CODE}
	}
	defer req.Body.Close()
	decoder := json.NewDecoder(req.Body)
	var youdao YouDao
	err = decoder.Decode(&youdao)
	if err != nil {
		fmt.Println(err)
		return *&YouDao{ErrorCode: ERROR_CODE}
	}
	return youdao
}
