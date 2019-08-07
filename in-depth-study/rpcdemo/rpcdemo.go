package rpcdemo

// rpc demo

import(
	"errors"
)


type DemoService struct {

}


type Args struct {
	A, B int
}


// 也可以有其他的结构体
type Args2 struct {
	S string
}



// RPC 的固定两个参数 第一个也可以是指针 result 必须是指针, 固定返回格式
func (DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}

	*result = float64(args.A) / float64(args.B)
	return nil
}


func (DemoService) Add(args Args, result *float64) error {
	*result = float64(args.A) + float64(args.B)
	return nil
}


// 例子2
func (DemoService) Change(args Args2, result *string) error {
	*result = args.S + "@"
	return nil
}