package main


import (
	"bufio"
	"os"
	"net"
	"fmt"
	"strings"
	"strconv"
	"time"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	
	fmt.Print("请输入连接IP:")
	input.Scan()

	// ip := "127.0.0.1:1200"
	ip := input.Text()

	tcpAddr, err := net.ResolveTCPAddr("tcp4", ip)		// 转换为 TCPAddr
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)		// 建立TCP连接,获得 TCPConn
	checkError(err)
	defer conn.Close()
	fmt.Println("success connection!")

	go Reader(conn)	// 读取数据

	for {
		input.Scan()
		command := input.Text()
		if command == "exi" {
			fmt.Println("exit!")
			_, err = conn.Write([]byte(command))
			break
		} else if len(command) >= 3 && command[:3] == "get" {
			_, err = conn.Write([]byte(command))	
		} else{
			_, err = conn.Write([]byte("tex"+command))
		}
	}
}


func Reader(conn net.Conn) {
	readBuff := make([]byte, 2048)
	for {
		size, err := conn.Read(readBuff)
		checkError(err)

		buff := string(readBuff[:size])
		switch(buff[:5]){
		case "text:":
			fmt.Println(strings.TrimSpace(buff))
		default:
			getFile(conn, &buff)
		}
		readBuff = make([]byte, 2048)
	}
}


func getFile(conn net.Conn, buff *string) {
	file, err := os.Create(strconv.FormatInt(time.Now().Unix(), 10) + *buff)
	checkError(err)
	defer file.Close()

	readBuff := make([]byte, 2048)
	for {
		size, err := conn.Read(readBuff)
		checkError(err)

		*buff = strings.TrimSpace(string(readBuff[:size]))
		if strings.Contains(*buff, "@EOF") {
			_, err = file.Write([]byte( strings.Replace(*buff, "@EOF", "", 1) ) )
			checkError(err)	
			break
		} else if strings.Contains(*buff, "@404") {
			fmt.Println("not found 404!")
			return
		}
		_, err = file.Write([]byte(*buff))
		checkError(err)
	}
	fmt.Println("---read file end---")
}


func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}