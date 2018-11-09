package main

import (
	"encoding/json"

	"gopkg.in/natefinch/lumberjack.v2"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

func log(path string, msg string) {

	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["` + path + `"],
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase"
		}
	  }`)

	var cfg zap.Config
	//cfg.OutputPaths = []string{path}
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, _ := cfg.Build()

	defer logger.Sync()

	//for index := 0; index < 100000; index++ {
	logger.Info(msg)
	//}
}

// func main() {
// 	w := zapcore.AddSync(&lumberjack.Logger{
// 		Filename: "/Users/wangchuan/workspace/go/src/study-go/04/zap/log.log",
// 		MaxSize:  100000, // megabytes
// 		//MaxBackups: 3,
// 		MaxAge: 28, // days
// 	})
// 	core := zapcore.NewCore(
// 		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
// 		w,
// 		zap.InfoLevel,
// 	)
// 	logger := zap.New(core)

// 	defer logger.Sync()

// 	for index := 0; index < 100000; index++ {
// 		logger.Error("error")
// 	}

// }

func log3() {
	topicErrors := zapcore.AddSync(&lumberjack.Logger{
		Filename: "/Users/wangchuan/workspace/go/src/study-go/04/zap/error.log",
	})
	infoErrors := zapcore.AddSync(&lumberjack.Logger{
		Filename: "/Users/wangchuan/workspace/go/src/study-go/04/zap/info.log",
	})
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, topicErrors, zap.ErrorLevel),
		zapcore.NewCore(encoder, infoErrors, zap.InfoLevel),
	)

	// From a zapcore.Core, it's easy to construct a Logger.
	logger := zap.New(core)
	defer logger.Sync()

	for index := 0; index < 100000; index++ {
		logger.Error("error")
	}

	for index := 0; index < 100000; index++ {
		logger.Info("error")
	}
}

func main() {
	// for index := 0; index < 100000; index++ {
	// 	log("/Users/wangchuan/workspace/go/src/study-go/04/zap/log1.log", "eee")
	// }

	//log("/Users/wangchuan/workspace/go/src/study-go/04/zap/log4.log", "eee")

	log3()
}
