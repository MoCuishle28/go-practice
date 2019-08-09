package worker


// 为了在 RPC 中传输函数信息
type SerializedParser struct {
	FuncionName string 		// 函数名
	Args interface{}		// 函数参数
}