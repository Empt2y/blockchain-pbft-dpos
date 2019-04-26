package core

const (
	InsDataTypeNormal = iota
	InsDataTypeTransaction
	InsDataTypeBlock
)

// 指令
type RawInstruction struct {
	Hash    []byte // 指令哈希
	Address []byte // 操作人地址
	Sign    []byte // 签名
	Data    []byte // 指令内容
	Type    int
}
