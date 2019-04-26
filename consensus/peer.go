package consensus

import "net"

// Peer结构体
type Peer struct {
	Id         int      // ID
	RemotePort int      // 端口
	LocalId    int      // 本地ID
	Conn       net.Conn // 连接
}

//
func (p *Peer) Send(msg *Message) {

}

//
func (p *Peer) Close() {
	p.Conn.Close()
}
