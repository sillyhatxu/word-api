package dao

import (
	"encoding/json"
	"github.com/sillyhatxu/mysql-client"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"testing"
	"word-api/client/youdaoclient"
)

const (
	dataSourceName = `sillyhat:sillyhat@tcp(161.117.82.136:3306)/sillyhat_dt`
	maxIdleConns   = 5
	maxOpenConns   = 10
)

func TestReadFile(t *testing.T) {
	plan, _ := ioutil.ReadFile("word.json")
	var data interface{}
	err := json.Unmarshal(plan, &data)
	assert.Nil(t, err)
	log.Println(data)
}

func TestInsertWord(t *testing.T) {
	dbclient.InitialDBClient(dataSourceName, maxIdleConns, maxOpenConns)
	id, err := InsertWord("tendency")
	assert.Nil(t, err)
	assert.EqualValues(t, id, 1)
}

func TestUpdateWord(t *testing.T) {
	dbclient.InitialDBClient(dataSourceName, maxIdleConns, maxOpenConns)
	translation := []string{"趋势"}
	explains := []string{"n. 倾向，趋势；癖好"}
	webArray := make([]*youdaoclient.Web, 3)
	value1 := []string{"倾向", "趋向", "趋势"}
	value2 := []string{"居中趋势", "集中趋势", "集中趋"}
	value3 := []string{"不正之风", "不正派的作风"}
	webArray[0] = &youdaoclient.Web{Key: "tendency", Value: value1}
	webArray[1] = &youdaoclient.Web{Key: "Central Tendency", Value: value2}
	webArray[2] = &youdaoclient.Web{Key: "unhealthy tendency", Value: value3}
	id, err := UpdateWord(*&youdaoclient.YouDao{
		Translation: translation,
		Basic: *&youdaoclient.Basic{
			PhoneticUS: "'tɛndənsi",
			Phonetic:   "tend(ə)nsɪ",
			PhoneticUK: "tend(ə)nsɪ",
			Explains:   explains,
		},
		Query:     "tendency",
		ErrorCode: 0,
		Web:       webArray,
	})
	assert.Nil(t, err)
	assert.EqualValues(t, id, 1)
}

func TestCount(t *testing.T) {
	dbclient.InitialDBClient(dataSourceName, maxIdleConns, maxOpenConns)
	count1, err := Count("tendency")
	assert.Nil(t, err)
	assert.EqualValues(t, count1, 1)
	count2, err := Count("Tendency")
	assert.Nil(t, err)
	assert.EqualValues(t, count2, 1)
	count3, err := Count(" Ten dEncy ")
	assert.Nil(t, err)
	assert.EqualValues(t, count3, 1)
	count4, err := Count(" Gen dEncy ")
	assert.Nil(t, err)
	assert.EqualValues(t, count4, 0)
}
