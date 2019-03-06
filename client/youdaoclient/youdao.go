package youdaoclient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const ERROR_CODE = -1

type YouDao struct {
	Translation []string `json:"translation"`
	Basic       Basic    `json:"basic"`
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
	timeout := time.Duration(30 * time.Second) //超时时间50ms
	client := &http.Client{Timeout: timeout}
	reqest, err := http.NewRequest("GET", "http://fanyi.youdao.com/openapi.do?keyfrom=SillyHatYouDao&key=987724779&type=data&doctype=json&version=1.1&q="+word, nil)
	if err != nil {
		return *&YouDao{ErrorCode: ERROR_CODE}
	}
	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		return *&YouDao{ErrorCode: ERROR_CODE}
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return *&YouDao{ErrorCode: ERROR_CODE}
	}
	defer response.Body.Close()
	var youdao YouDao
	jsonErr := json.Unmarshal([]byte(string(body)), &youdao)
	if jsonErr != nil {
		return *&YouDao{ErrorCode: ERROR_CODE}
	}
	return youdao
}
