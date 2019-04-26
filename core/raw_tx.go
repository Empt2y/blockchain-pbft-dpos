package core

// 交易
type RawTransaction struct {
	Hash      []byte // 交易哈希
	From      []byte // 发送者
	To        []byte // 接收者
	Type      int32  // 交易类型
	Size      int32  // 交易大小
	Status    int32  // 交易状态
	Timestamp int64  // 时间戳
	Fee       int64  // 手续费
	Message   []byte // 执行指令内容
}
