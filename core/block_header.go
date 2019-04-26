package core

// 区块头
type BlockHeader struct {
	Version       int32    // 版本号
	ChainFlag     int32    // 链标识
	Number        int64    // 区块高度/区块编号
	Timestamp     int64    // 时间戳
	WitnessReward int64    // 见证人奖励
	Address       []byte   // 见证人地址
	PrevHash      []byte   // 前区块哈希
	MerkleRoot    []byte   // 交易根哈希
	Forks         [][]byte // 分叉区块集合
}
