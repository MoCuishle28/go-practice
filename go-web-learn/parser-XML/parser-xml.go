package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	// `...` 这个是struct的一个特性 被称为 struct tag 是用来辅助反射的
	// tag 的字段名、XML元素都是大小写敏感的
	XMLName     xml.Name `xml:"servers"`

	// tag定义了中含有 ",attr" 那么解析的时候就会将该结构所对应的 element 的与字段同名的属性的值赋值给该字段
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`			// 叶节点结构体

	// 如果 struct 的一个字段是 string 或者 []byte 类型且它的 tag 含有 ",innerxml"
	// Unmarshal 将会将此字段所对应的元素内所有内嵌的原始 xml 累加到此字段上
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func main() {
	file, err := os.Open("servers.xml") // For read access.		
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)		// 解析字节切片 data 存储在 v 中
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
	fmt.Println("-------")

	fmt.Println(v.XMLName)
	fmt.Println(v.Version)
	for k,v := range v.Svs{
		fmt.Println(k, v)
	}
	fmt.Println(v.Description)
}