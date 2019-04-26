package core

import (
	"crypto/sha256"
)

// 区块
type Block struct {
	Hash   []byte       // 区块哈希
	Header *BlockHeader // 区块头
	Body   *BlockBody   // 区块体
}

// 创建创世块
func CreateGenesis() *Block {
	b := &Block{}
	b.CalcHash()
	return b
}

// 计算区块哈希
func (b *Block) CalcHash() {
	s := ""
	h := sha256.New()
	h.Write([]byte(s))
	b.Hash = h.Sum(nil)
}
