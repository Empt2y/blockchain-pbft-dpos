package consensus

// 节点
type Node struct {
	Id      int32
	IsBad   bool
	PeerIds []int32
	Peers   []Peer
}
