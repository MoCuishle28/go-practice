package main


import(
	"net/http"
	"os"
	"io/ioutil"
	"log"
	_ "net/http/pprof"	// 用于服务器性能分析  访问 http://127.0.0.1:8888/debug/pprof/ 可以查看性能情况\
	//  还可以用查看耗时操作 go tool pprof http://127.0.0.1:8888/debug/profile
)

/*
统一错误处理
*/

// 定义新的类型, 处理业务逻辑的函数，返回一个error
type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 返回标准的实现 http/net 里面被路由选择的处理函数, 同一大类的错误统一处理
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	// 返回的函数才是实现接口的可以被路由选择的函数
	return func(writer http.ResponseWriter, request *http.Request) {
		// recover处理业务逻辑中可能出现的其他运行中被panic的错误
		// panic 之后会寻找是否遇到recover，若没有则退出程序（web服务受保护不会直接退出，会进入系统写的recover）
		defer func() {
			// recover
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, 
					http.StatusText(http.StatusInternalServerError), 
					http.StatusInternalServerError)
			}
		} ()

		// 在返回函数内部调用对应的业务逻辑函数
		err := handler(writer, request)
		if err != nil {
			log.Printf("Error: %v", err)
			code := http.StatusOK
			switch {
				case os.IsNotExist(err):
					code = http.StatusNotFound
				default:
					code = http.StatusInternalServerError	// 500错误(不知道什么错误)
			}
			// 第一个是响应，第二个返回的字符串，第三个是错误代码
			http.Error(writer, http.StatusText(code), code)
		}
	}
}


// 这个不再是直接被 handleFunc 关联到的函数了
func getFile(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/") : ]	// 拿到url的后面部分，输入时输入一个文件名字
	file, err := os.Open(path)
	if err != nil {
		// http.Error(writer, err.Error(), http.StatusInternalServerError)
		// 统一处理错误后，遇到错误直接返回就行
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		// panic(err)
		return err
	}
	writer.Write(all)
	return err
}


func main() {
	// 用errWrapper包装一下
	http.HandleFunc("/list/", errWrapper(getFile))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}