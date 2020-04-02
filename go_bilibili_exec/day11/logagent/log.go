package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func convertLogLevel(level string) int {
	var logLevel int
	switch level {
	case "debug":
		logLevel = logs.LevelDebug
	case "warn":
		logLevel = logs.LevelWarn
	case "info":
		logLevel = logs.LevelInfo
	case "trace":
		logLevel = logs.LevelTrace
	}

	return logLevel
}

func initLogger() (err error) {
	config := make(map[string]interface{})
	config["filename"] = appConfig.LogPath
	config["level"] = convertLogLevel(appConfig.LogLevel)

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("initLogger failed, marshal failed, err:", err)
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}