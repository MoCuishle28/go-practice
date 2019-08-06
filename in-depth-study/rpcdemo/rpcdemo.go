package rpcdemo


import(
	"errors"
)


type DemoService struct {

}


type Args struct {
	A, B int
}



// RPC 的固定两个参数 第一个也可以是指针 result 必须是指针, 固定返回格式
func (DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}

	*result = float64(args.A) / float64(args.B)
	return nil
}