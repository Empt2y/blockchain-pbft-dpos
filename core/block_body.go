package core

// 区块体
type BlockBody struct {
	Txs []*RawTransaction
}

// 计算区块体哈希
func (bb *BlockBody) CalcHash() {

}
