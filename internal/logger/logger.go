package logger

import (
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/config"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	log *logrus.Logger
	cfg *config.Config
}

func New(cfg *config.Config) *Logger {
	log := logrus.New()
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true

	switch cfg.LogLevel {
	case "DEBUG":
		log.SetLevel(logrus.DebugLevel)
	case "INFO":
		log.SetLevel(logrus.InfoLevel)
	case "ERROR":
		log.SetLevel(logrus.ErrorLevel)
	case "FATAL":
		log.SetLevel(logrus.FatalLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}
	return &Logger{
		cfg: cfg,
		log: log,
	}
}

func (l *Logger) LogInfo(message string, args ...interface{}) {
	l.log.Infoln(message, args)
}

func (l *Logger) LogError(message string, args ...interface{}) {
	l.log.Error(message, args)
}

func (l *Logger) LogFatal(message string, args ...interface{}) {
	l.log.Fatalln(message, args)
}
