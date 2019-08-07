package main

import(
	"testing"
	"time"

	"Go-practice/in-depth-study/reptile-distributed/rpcsupport"
	"Go-practice/in-depth-study/reptile-project/model"
)


func TestItemSaver(t *testing.T) {
	const host = ":1234"
	// start ItemSaverServer
	go serveRpc(host, "test1")

	time.Sleep(8*time.Second)

	// start ItemSaverClient
 	client, err := rpcsupport.NewClient(host)
 	if err != nil {
 		panic(err)
 	}

	// call save
	profile := model.Profile{
		Name: "aaa",
		Gender: "男",
		Age: 25,
		Height: 175,
		Weight: 60,
		Income: "50000-80000",
		Marriage: "未婚",
		Education: "硕士",
		Occupation: "程序员",
		Hokou: "广东深圳",
		House: "已购房",
		Car:	"已购车",
	}
	result := ""
	err = client.Call("ItemSaverService.Save", profile, &result)
	if err != nil || result == "ok" {
		t.Errorf("item:%+v result: %s; err:%v", profile, result, err)
	}
}