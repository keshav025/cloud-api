package logs

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func Init(isLocalEnv bool) {
	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
	})

	if isLocalEnv {
		logrus.SetFormatter(&prefixed.TextFormatter{
			ForceFormatting: true,
		})
	}

	logrus.SetLevel(logrus.DebugLevel)
}
