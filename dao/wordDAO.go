package dao

import (
	"encoding/json"
	"github.com/sillyhatxu/mysql-client"
	"word-api/client/youdaoclient"
)

const (
	insert_sql = `
		INSERT INTO youdao_word 
		(word, translation, phonetic, us_phonetic, uk_phonetic, explains, web, status, created_date, last_modified_date)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, now(), now())
	`
	update_sql = `
		UPDATE youdao_word
		SET translation        = ?,
		    phonetic           = ?,
		    us_phonetic        = ?,
		    uk_phonetic        = ?,
		    explains           = ?,
		    web                = ?,
		    status             = 1,
		    last_modified_date = now()
		WHERE upper(replace(word,' ','')) = upper(replace(?,' ',''))
	`
	count_sql = `
		select count(1) from youdao_word where upper(replace(word,' ','')) = 
	`
)

func InsertWord(word string) (int64, error) {
	id, err := dbclient.Client.Insert(insert_sql, word, "", "", "", "", "", "", 0)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func InsertWordDetail(youdao youdaoclient.YouDao) (int64, error) {
	translation, err := json.Marshal(youdao.Translation)
	if err != nil {
		return 0, err
	}
	explains, err := json.Marshal(youdao.Basic.Explains)
	if err != nil {
		return 0, err
	}
	web, err := json.Marshal(youdao.Web)
	if err != nil {
		return 0, err
	}
	id, err := dbclient.Client.Insert(insert_sql, youdao.Query,
		translation,
		youdao.Basic.Phonetic,
		youdao.Basic.PhoneticUS,
		youdao.Basic.PhoneticUK,
		explains,
		web, 1)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func UpdateWord(youdao youdaoclient.YouDao) (int64, error) {
	translation, err := json.Marshal(youdao.Translation)
	if err != nil {
		return 0, err
	}
	explains, err := json.Marshal(youdao.Basic.Explains)
	if err != nil {
		return 0, err
	}
	web, err := json.Marshal(youdao.Web)
	if err != nil {
		return 0, err
	}
	id, err := dbclient.Client.Update(update_sql,
		translation,
		youdao.Basic.Phonetic,
		youdao.Basic.PhoneticUS,
		youdao.Basic.PhoneticUK,
		explains,
		web,
		youdao.Query)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func Count(word string) (int64, error) {
	sql := count_sql + " upper(replace('" + word + "',' ',''))"
	count, err := dbclient.Client.Count(sql)
	if err != nil {
		return 0, err
	}
	return count, nil
}
