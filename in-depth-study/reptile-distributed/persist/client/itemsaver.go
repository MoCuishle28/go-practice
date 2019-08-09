package client


import(
	"log"

	"Go-practice/in-depth-study/reptile-distributed/rpcsupport"
)


func ItemSaver(host string) chan interface{} {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil
	}

	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out

			// 每收到一个元素 call rpc save item
			result := ""
			err := client.Call("ItemSaverService.Save", item, &result)
			if err != nil || result != "ok" {
				log.Printf("saver error:item #%d: %v; error:%v\n", itemCount, item, err)
				continue
			}
			log.Printf("save item #%d: %v\n", itemCount, item)
			itemCount++
		}
	} ()
	return out
}