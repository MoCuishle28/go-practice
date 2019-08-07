package persist


import(
	// "log"
	// "Go-practice/in-depth-study/reptile-project/engine"
)


type ItemSaverService struct{
	// Client *elastic.Client		// 用于发送指令
	Index string 				// 在 elasticsearch 中类似于库名
}


func (s *ItemSaverService) Save(item interface{}, result *string) error {
	// log.Printf("Item saver: %v\n", item)
	// 存入 elassearch TODO
	*result = "ok"
	return nil
}