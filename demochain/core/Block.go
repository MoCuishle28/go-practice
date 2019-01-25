package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct{
	Index int64  			// 区块编号
	Timestamp int64			//时间戳
	PrevBlockHash string	//上一个区块的hash
	Hash string				// 当前区块hash

	// 简便起见 区块头、区块体放在同一个结构体
	Data string				//区块体数据(简易表示为字符串吧...)
}

func calculateHash(b Block) string {
	blockData := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash + b.Data
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}

// 产生新区块
func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.PrevBlockHash = preBlock.Hash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

// 生成创始区块
func GenerateGenesisBlock() Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	return GenerateNewBlock(preBlock, "Genesis Block")
}