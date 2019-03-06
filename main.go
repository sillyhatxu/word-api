package main

import (
	"encoding/json"
	"github.com/sillyhatxu/mysql-client"
	"io/ioutil"
	"log"
	"time"
	"word-api/client/youdaoclient"
	"word-api/dao"
)

//func init() {
//	cfgFile := flag.String("c", "config.conf", "configuration file")
//	flag.Parse()
//	config.ParseConfig(*cfgFile)
//}

type Word struct {
	Title string   `json:"title"`
	Words []string `json:"words"`
}

const (
	dataSourceName = `sillyhat:sillyhat@tcp(161.117.82.136:3306)/sillyhat`
	maxIdleConns   = 5
	maxOpenConns   = 10
)

func main() {
	dbclient.InitialDBClient(dataSourceName, maxIdleConns, maxOpenConns)
	wordJson, _ := ioutil.ReadFile("word.json")
	var wordArray []Word
	err := json.Unmarshal([]byte(string(wordJson)), &wordArray)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(wordArray); i++ {
		words := wordArray[i]
		for j := 0; j < len(words.Words); j++ {
			word := words.Words[j]
			log.Println("----- sleep -----")
			time.Sleep(5 * time.Second)
			log.Println("select word : " + word)
			count, err := dao.Count(word)
			if err != nil {
				continue
			}
			if count > 0 {
				continue
			}
			youdao := youdaoclient.Translation(word)
			if youdao.ErrorCode != youdaoclient.ERROR_CODE {
				_, err := dao.InsertWordDetail(youdao)
				if err != nil {
					continue
				}
			} else {
				_, err := dao.InsertWord(word)
				if err != nil {
					continue
				}
			}
		}
		//var data interface{}
		//err := json.Unmarshal(plan, &data)
		//assert.Nil(t, err)
		//log.Println(data)
		//dep ensure
		//logFormatter := &logrus.JSONFormatter{
		//	FieldMap: *&logrus.FieldMap{
		//		logrus.FieldKeyMsg: "message",
		//	},
		//}
		//conn, err := net.Dial("tcp", "localhost:51401")
		//if err != nil {
		//	log.Fatal("net.Dial error.", err)
		//}
		//hook := log.New(conn, logFormatter)
		//logrusConfig := log.NewLogrusConfig(logFormatter, logrus.DebugLevel, logrus.Fields{"module": "word-api"}, true, hook)
		//err = logrusConfig.InstallConfig()
		//if err != nil {
		//	log.Fatal("logrus config initial error.", err)
		//}
		//dbclient.InitialDBClient(config.Conf.MysqlDB.DataSource, config.Conf.MysqlDB.MaxIdleConns, config.Conf.MysqlDB.MaxOpenConns)
		//go scheduler.InitialScheduler(AutoQuery, "00:00:00", "24h")

		//api.InitialAPI()
	}
}
