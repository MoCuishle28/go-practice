package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func main() {
	config := make(map[string]interface{})
	config["filename"] = "./logcollect.log"
	// config["level"] = logs.LevelDebug		// 能写入的日志的级别，Debug以及以上的级别能写入
	config["level"] = logs.LevelInfo			// 这样就不能写入debug级别的日志了

	configStr, err := json.Marshal(config)	//返回一个byte数组
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))		// 以json字符串传进配置

	logs.Debug("this is a test, my name is %s", "stu01")
	logs.Trace("this is a trace, my name is %s", "stu02")
	logs.Warn("this is a warn, my name is %s", "stu03")
}
