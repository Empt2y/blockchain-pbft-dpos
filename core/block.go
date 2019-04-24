package core

// 区块
type Block struct {
	Height    int64
	Timestamp int64
	Hash      []byte
	Prev      []byte
	Txs       []*Transaction
}

// 创建创世块
func CreateGenesis() *Block {
	b := &Block{}
	b.SetHash()
	return b
}

// 计算哈希
func (b *Block) SetHash() {

}
