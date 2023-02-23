package canonicallog

import (
	"fmt"
	"testing"

	"github.com/multiformats/go-multiaddr"
	"github.com/yspk/go-libp2p-core/test"
)

func TestLogs(t *testing.T) {
	LogMisbehavingPeer(test.RandPeerIDFatal(t), multiaddr.StringCast("/ip4/1.2.3.4"), "somecomponent", fmt.Errorf("something"), "hi")
	LogMisbehavingPeerNetAddr(test.RandPeerIDFatal(t), dummyNetAddr{}, "somecomponent", fmt.Errorf("something"), "hi")
}

type dummyNetAddr struct{}

func (d dummyNetAddr) Network() string { return "tcp" }
func (d dummyNetAddr) String() string  { return "127.0.0.1:80" }
