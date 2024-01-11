package log

import (
	timeFormat "DistroJudge/time"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"runtime"
	"time"
)

// level
const (
	SLOW = iota
	INFO
	WARN
	ERROR
)

type DbLogConfig struct {
	ServiceName   string `yaml:"service-name"`
	Level         int64  `yaml:"level"`
	Mongo         bool   `yaml:"mongo"`
	MongoURI      string `yaml:"mongo-URI"`
	LogDatabase   string `yaml:"log-database"`
	LogCollection string `yaml:"log-collection"`
}

type DbLog struct {
	ctx           context.Context
	Client        *mongo.Client
	config        *DbLogConfig
	ServiceName   string
	level         int64
	mongo         bool
	logDatabase   string
	logCollection string
}

var DLog *DbLog

func NewClient(c *DbLogConfig) *DbLog {
	DLog = &DbLog{
		ctx:           context.Background(),
		ServiceName:   c.ServiceName,
		config:        c,
		level:         c.Level,
		logCollection: c.LogCollection,
		logDatabase:   c.LogDatabase,
		mongo:         c.Mongo,
	}

	var err error
	if c.Mongo {
		DLog.Client, err = mongo.Connect(DLog.ctx, options.Client().ApplyURI(c.MongoURI))
		if err != nil {
			panic(err)
		}
	}
	return DLog
}

func Infof(message string, v ...interface{}) {
	log(INFO, "info", fmt.Sprintf(message, v...))
}

func Warnf(message string, v ...interface{}) {
	log(WARN, "warn", fmt.Sprintf(message, v...))
}

func Errorf(message string, v ...interface{}) {
	log(ERROR, "error", fmt.Sprintf(message, v...))
}

func log(level int, levelName, message string) {
	if DLog.level < int64(level) {
		return
	}

	//fmt.Println(message)
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("yxlog Recovered err:%+v", err)
		}
	}()
	// 发送数据
	_, file, line, _ := runtime.Caller(2)
	strLog := fmt.Sprintf("## %s ## %s ## %s ## %s ## %s:%d",
		time.Now().Format(timeFormat.TimeLayout),
		levelName,
		DLog.ServiceName,
		message,
		file,
		line)

	if DLog.Client == nil {
		NewClient(DLog.config)
	}

	fmt.Println(strLog)
	if !DLog.mongo {
		return
	}

	m := bson.M{
		"time":    time.Now().Format(timeFormat.TimeLayout),
		"log":     message,
		"service": DLog.ServiceName,
		"level":   level,
	}
	_, err := DLog.Client.Database(DLog.logDatabase).Collection(DLog.logCollection).InsertOne(DLog.ctx, m)
	//_, err := DLog.Write([]byte(strLog))
	if err != nil {
		//超时失败重连
		NewClient(DLog.config)
	}
}

func (l *DbLog) Close() {
	_ = l.Client.Disconnect(l.ctx)
}
