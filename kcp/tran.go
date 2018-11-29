package kcp

import (
	"net"

	"github.com/jeckbjy/fairy"
	"github.com/jeckbjy/fairy/base"
	kcpgo "github.com/xtaci/kcp-go"
)

// NewTran create kcp transport
func NewTran() fairy.ITran {
	return base.NewTran(&kcpTran{})
}

// TODO: support option
type kcpTran struct {
}

func (kt *kcpTran) Connect(host string) (net.Conn, error) {
	return kcpgo.DialWithOptions(host, nil, 10, 3)
}

func (kt *kcpTran) Listen(host string) (net.Listener, error) {
	return kcpgo.ListenWithOptions(host, nil, 10, 3)
}

func (kt *kcpTran) Serve(l net.Listener, cb base.OnAccept) {
	for {
		conn, err := l.Accept()
		cb(conn, err)
	}
}
