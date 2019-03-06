package main

import (
	"flag"
	log "github.com/sillyhatxu/microlog"
	"github.com/sillyhatxu/mysql-client"
	"github.com/sirupsen/logrus"
	"net"
	"word-api/api"
	"word-api/config"
)

func init() {
	cfgFile := flag.String("c", "config.conf", "configuration file")
	flag.Parse()
	config.ParseConfig(*cfgFile)
}

func main() {
	//dep ensure
	logFormatter := &logrus.JSONFormatter{
		FieldMap: *&logrus.FieldMap{
			logrus.FieldKeyMsg: "message",
		},
	}
	conn, err := net.Dial("tcp", "localhost:51401")
	if err != nil {
		log.Fatal("net.Dial error.", err)
	}
	hook := log.New(conn, logFormatter)
	logrusConfig := log.NewLogrusConfig(logFormatter, logrus.DebugLevel, logrus.Fields{"module": "word-api"}, true, hook)
	err = logrusConfig.InstallConfig()
	if err != nil {
		log.Fatal("logrus config initial error.", err)
	}
	dbclient.InitialDBClient(config.Conf.MysqlDB.DataSource, config.Conf.MysqlDB.MaxIdleConns, config.Conf.MysqlDB.MaxOpenConns)
	api.InitialAPI()
	log.Info("---------- Project Close ----------")
}
