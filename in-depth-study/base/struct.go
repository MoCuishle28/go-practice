package main


import(
	"fmt"
)


type point struct {
	i, j int
}


// 对象方法
func (p point) sum() int {
	return p.i + p.j
}


// 这里的p依然是传值
// 甚至等于nil的point也能调用这些函数，因为只是把point类型的变量传进函数而已，nil也一样传进来了。
func (p point) setI() {
	p.i = 100
}

// 改成指针接收
func (p *point) setI_by_ref() {
	if p == nil {
		fmt.Println("is nil !!!")
		return
	}
	p.i = 100		// 不需要加*，可以直接使用
}


func f(p point) {
	p.i = 10
	fmt.Println("f:", p)
}

func f0(p *point) {
	p.i = 10		// 这样也可以直接使用，不需要加*
	fmt.Println("f0:", p)
}


// go可以使用工厂函数创建对象
func createPoint(i, j int) *point {
	// 返回地址会在堆上分配，否则在栈上分配？ 栈上分配在函数结束时会被回收
	return &point{i:i, j:j}			// 可以返回局部变量的地址
}


func main() {
	p := point{i:3, j:4}	// 一个point类型的值
	fmt.Println(p)
	f(p)
	fmt.Println(p)

	p0 := &point{i:3, j:4}	// 一个point类型的指针
	fmt.Println(p0)
	f0(p0)
	fmt.Println(p0)

	pp := createPoint(3, 4)
	fmt.Println("after createPoint:", pp)
	fmt.Println("p.sum():", p.sum())

	p.setI()
	fmt.Println("p.setI():", p.i)
	p.setI_by_ref()
	fmt.Println("p.setI_by_ref():", p.i)

	var po *point
	fmt.Println(po)
	po.setI_by_ref()
}