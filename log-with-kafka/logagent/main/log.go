package main

import(
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)


func convertLogLevel(level string) int {
	switch(level) {
		case "debug":
			return logs.LevelDebug
		case "warn":
			return logs.LevelWarn
		case "info":
			return logs.LevelInfo
		case "trace":
			return logs.LevelTrace
	}
	return logs.LevelDebug
}


func initLogger() (err error) {

	config := make(map[string]interface{})
	config["filename"] = appConfig.logPath
	config["level"] = convertLogLevel(appConfig.logLevel)

	configStr, err := json.Marshal(config)	//返回一个byte数组
	if err != nil {
		fmt.Println("initLogger failed, marshal err:", err)
		return err
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))		// 以json字符串传进配置
	return
}