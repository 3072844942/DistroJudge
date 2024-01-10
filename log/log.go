package log

import (
	timeFormat "DistroJudge/time"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"runtime"
	"time"
)

// level
const (
	SLOW  = 0
	INFO  = 1
	WARN  = 2
	ERROR = 3
)

type DbLogConfig struct {
	ServiceName   string `yaml:"service-name"`
	MongoURI      string `yaml:"mongo-URI"`
	Level         int64  `yaml:"level"`
	LogDatabase   string `yaml:"log-database"`
	LogCollection string `yaml:"log-collection"`
}

type DbLog struct {
	Client        *mongo.Client
	config        *DbLogConfig
	ServiceName   string
	level         int64
	logDatabase   string
	logCollection string
}

var DLog *DbLog

func NewClient(c *DbLogConfig) *DbLog {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(c.MongoURI))
	if err != nil {
		Errorf("failed to connect mongo. err: %v", err)
	}

	DLog = &DbLog{
		ServiceName:   c.ServiceName,
		config:        c,
		Client:        client,
		level:         c.Level,
		logCollection: c.LogCollection,
		logDatabase:   c.LogDatabase,
	}
	return DLog
}

func Infof(message string, v ...interface{}) {
	log("info", fmt.Sprintf(message, v...))
}

func Warnf(message string, v ...interface{}) {
	log("warn", fmt.Sprintf(message, v...))
}

func Errorf(message string, v ...interface{}) {
	log("error", fmt.Sprintf(message, v...))
}

func log(level, message string) {
	//fmt.Println(message)
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("yxlog Recovered err:%+v", err)
		}
	}()
	// 发送数据
	_, file, line, _ := runtime.Caller(2)
	strLog := fmt.Sprintf("## %s ## %s ## %s ## action ## %s ## %s:%d",
		time.Now().Format(timeFormat.TimeLayout),
		level,
		DLog.ServiceName,
		message,
		file,
		line)

	if DLog.Client == nil {
		NewClient(DLog.config)
	}

	_, err := DLog.Write([]byte(strLog))
	if err != nil {
		//超时失败重连
		NewClient(DLog.config)
	}

}

func (l *DbLog) Close() {
	_ = l.Client.Disconnect(context.Background())
}

func (l *DbLog) Write(b []byte) (n int, err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("dlog Recovered err: %+v", err)
		}
	}()

	_, err = l.Client.Database(l.logDatabase).Collection(l.logCollection).InsertOne(context.Background(), b)
	if err != nil {
		Warnf("dbLog err:%+v", err)
	}
	return
}
