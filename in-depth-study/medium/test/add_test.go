package main


import(
	"testing"
)

/*
测试文件名字格式要是 *_test.go
测试函数名字格式要是  func Test*(t *testing.T) { ... }

进入目录下 go test . 可以自动运行测试

查看代码覆盖率：
	将覆盖率信息写入 c.out 文件
	go test -coverprofile=c.out

	打开一个 html 可视化覆盖(go tool 工具还有很多功能)
	go tool cover -html=c.out
	绿色的是覆盖的代码
	红色是未覆盖的代码

性能测试：（执行 func Benchmark*(b *testing.B) { ... } 函数）
	前提是功能测试函数 Test* 不能出错
	go test -bench .
	BenchmarkAdd-4  2000000000 0.34 ns/op  指一共执行2000000000次，平均每条执行时间 0.34 ns

	需要先装一个工具
	生成一个 cpu.out
	go test -bench . -cpuprofile cpu.out
	查看具体性能信息
	go tool pprof cpu.out
*/


// 面向表单测试 (是一种编写测试用例的思路)
// 允许有个别不通过的测试用例，可以专注于测试数据设计
func TestAdd(t *testing.T) {
	// 初始化一个测试数据表 (匿名的struct)
	tests := [] struct { a, b, c int } {
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},		// 故意写错一个测试用例 看看结果与用例结果不一致的状况
		{30000, 40000, 50000},
	}

	for _, tt := range tests {
		// 先算出来actual 再判断是否为正确答案
		if actual := cal(tt.a, tt.b); actual != tt.c {
			t.Errorf("cal(%d, %d); got %d expected %d\n", tt.a, tt.b, actual, tt.c)
		}
	}
}


// 性能测试
func BenchmarkAdd(b *testing.B) {
	// 找一个运算量大的
	a, bb, c := 342, 1520, 1558

	// b.N是让编译器决定重复运算多少次
	for i := 0; i < b.N; i++ {
		if actual := cal(a, bb); actual != c {
			b.Errorf("cal(%d, %d); got %d expected %d\n", a, bb, actual, c)
		}
	}
}