package main


import(
	"fmt"
)


type TestStruct struct {
	num int
	str string
}


// toString()方法
func (ts *TestStruct) String() string {
	return fmt.Sprintf("(num:%d, str:%s)\n", ts.num, ts.str)
}

// 还有Reader/Writer接口


func main() {
	ts := TestStruct{num:99, str:"aaa"}
	fmt.Println(ts)
}