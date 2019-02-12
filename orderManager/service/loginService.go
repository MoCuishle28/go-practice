package service

import (
	// 打印
	"fmt"

	// 用于文件操作
	"os"
	"encoding/json"
	"io/ioutil"
)


// 读取 root.json 信息
type RootLogin struct {
	Username string
	Password string
}


func ValidLogin(username, pwd string) int {
	var root_login RootLogin
	var ret int
	status := map[string]int{"success":1, "name_err":2, "pwd_err":3, "err":-1}
	file, err := os.Open("root.json")	// 用相对路径读取 文件放在 main 包中
	if err != nil {
		fmt.Println(err)
		return status["err"]
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return status["err"]
	}

	json.Unmarshal(data, &root_login)	// 解析 JSON
	root_name := root_login.Username
	root_pwd := root_login.Password

	if root_name == username {
		if root_pwd == pwd {
			ret = status["success"]
		} else {
			ret = status["pwd_err"]
		}
	} else {
		ret = status["name_err"]
	}
	return ret
}