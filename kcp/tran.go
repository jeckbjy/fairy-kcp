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

func (kt *kcpTran) Connect(host string, options ...fairy.Option) (net.Conn, error) {
	var kop *KcpOption
	for _, op := range options {
		if t, ok := op.(*KcpOption); ok {
			kop = t
			break
		}
	}
	if kop != nil {
		return kcpgo.DialWithOptions(kop.Addr, kop.Block, kop.DataShards, kop.ParityShards)
	} else {
		return kcpgo.Dial(host)
	}
}

func (kt *kcpTran) Listen(host string, options ...fairy.Option) (net.Listener, error) {
	var kop *KcpOption
	for _, op := range options {
		if t, ok := op.(*KcpOption); ok {
			kop = t
			break
		}
	}
	if kop != nil {
		return kcpgo.ListenWithOptions(kop.Addr, kop.Block, kop.DataShards, kop.ParityShards)
	} else {
		// return kcpgo.ListenWithOptions(host, nil, 10, 3)
		return kcpgo.Listen(host)
	}
}

func (kt *kcpTran) Serve(l net.Listener, cb base.OnAccept) {
	for {
		conn, err := l.Accept()
		cb(conn, err)
	}
}

// KcpOption 配置选项
type KcpOption struct {
	Addr         string
	Block        kcpgo.BlockCrypt
	DataShards   int
	ParityShards int
}

// WithKcpOption 创建KcpOption
func WithKcpOption(addr string, block kcpgo.BlockCrypt, dataShards int, parityShards int) *KcpOption {
	return &KcpOption{Addr: addr, Block: block, DataShards: dataShards, ParityShards: parityShards}
}
