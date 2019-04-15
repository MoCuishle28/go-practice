package main


import (
	"fmt"
	"net"
	"os"
	"io/ioutil"
	"time"
	"strings"
)


func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	defer listener.Close()

 	var userID int64 = 0
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("Accept user:", userID, "...")
		go handleClient(conn, userID)
		userID++
	}
}


func handleClient(conn net.Conn, userID int64) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
	defer conn.Close()

	dir_list, e := ioutil.ReadDir("I:\\Go_WorkSpace\\src\\Go-practice\\network-experiment\\experiment1\\server")
	checkError(e)

	list := make([]string, 1)
	list = append(list, "text:")
	for _, v := range dir_list {
		list = append(list, v.Name())
	}
	dir := strings.TrimSpace(strings.Join(list, "\n"))
	conn.Write([]byte(dir))

	request := make([]byte, 128)
	for {
		size, err := conn.Read(request)		// 读取客户发来的信息到 request 中
		checkError(err)

		command := string(request[:size])
		switch(command[:3]) {
		case "get":
			fmt.Println("ID:", userID, " ", strings.TrimSpace(command))
			filename := strings.Split(command, " ")[1]
			WriteFile(conn, filename)
		case "exi":
			fmt.Println("ID:", userID, " Close...")
			return
		case "tex": 
			fmt.Println("ID:", userID, " ", strings.TrimSpace("return text:"+strings.ToUpper(command[3:])))
			conn.Write([]byte(strings.TrimSpace("text:"+strings.ToUpper(command[3:]))))
		}
		request = make([]byte, 128)
	}
}


func WriteFile(conn net.Conn, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("not found! ", filename)
		conn.Write([]byte("@404"))
		return
	}
	// checkError(err)
	defer file.Close()

	buff := make([]byte, 2048)

	conn.Write([]byte(filename))		// 先发送文件名字
	for {
		size, _ := file.Read(buff)
		if size == 0 {
			conn.Write([]byte("@EOF"))
			break
		}
		conn.Write(buff[:size])
	}
	fmt.Println("---"+filename+" send end---")
}


func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}