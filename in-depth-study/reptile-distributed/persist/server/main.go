package main


// 提供 RPC 服务
import(
	"log"

	"Go-practice/in-depth-study/reptile-distributed/rpcsupport"
	"Go-practice/in-depth-study/reptile-distributed/persist"
)


func main() {
	// 出错则强制退出
	log.Fatal(serveRpc(":1234", "zhenai"))
}


func serveRpc(host, index string) error {
	// 还应该先创建用于 elasticsearch 操作的 client
	return rpcsupport.ServeRpc(":1234", &persist.ItemSaverService{
		Index: index,
	})
}