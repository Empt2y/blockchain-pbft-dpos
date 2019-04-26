package consensus

// 消息类型
const (
	MsgTypeInit = iota + 1
	MsgTypeBlock
	MsgPrepare
	MsgCommit
)

// 消息结构体
type Message struct {
	MsgType int // 消息类型
	Id      int // 消息ID
}
