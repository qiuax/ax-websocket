package logger

import (
	"ax-websocket/conf"
	"ax-websocket/helper"
	"log"

	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongodbHooker struct {
	Session *mgo.Session
}

func (h *mongodbHooker) Fire(entry *logrus.Entry) error {
	data := bson.M(GetLogrusData(entry))

	go func() {
		s := h.Session.Copy()
		defer s.Close()
		mgoErr := s.DB(conf.MongoLogsDatabase).C(helper.FormatDate("Ymd", 0)).Insert(data)
		if mgoErr != nil {
			log.Printf("Failed to send log entry to mongodb: %v", mgoErr)
		}
	}()

	return nil
}

func (h *mongodbHooker) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

func GetLogrusData(entry *logrus.Entry) map[string]interface{} {
	data := make(map[string]interface{})
	data["level"] = entry.Level.String()
	//data["time"] = entry.Time.Format(time.RFC3339)
	data["msg"] = entry.Message

	for k, v := range entry.Data {
		if errData, isError := v.(error); logrus.ErrorKey == k && v != nil && isError {
			data[k] = errData.Error()
		} else {
			data[k] = v
		}
	}

	return data
}
