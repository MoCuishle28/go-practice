package main

import (
	"net/http"
	"Go-practice/demochain/core"
	"io"
	"encoding/json"
)

var blockchain *core.Blockchain

func run() {
	// 参数为：1.访问路径 2.处理访问的函数
	http.HandleFunc("/blockchain/get", blockchainGetHandle)	//获取区块
	http.HandleFunc("/blockchain/write", blockchainWriteHandle)	//写入区块
	http.ListenAndServe("localhost:8888", nil)
}

func blockchainGetHandle(w http.ResponseWriter, r *http.Request) {
	// 返回 字节切片，error
	bytes, error := json.Marshal(blockchain)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))	// 这样就可以返回
}

func blockchainWriteHandle(w http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("data")
	blockchain.SendData(blockData)
	blockchainGetHandle(w, r)
}

func main() {
	blockchain = core.NewBlockchain()
	run()
}