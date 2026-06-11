package node

import (
	"context"
	"encoding/json"

	"github.com/libp2p/go-libp2p/core/peer"
	libp2pproto "github.com/libp2p/go-libp2p/core/protocol"
	"github.com/neuroroot/core/pkg/acp"
)

// initACP يهيئ بروتوكول ACP على العقدة
func (n *Node) initACP() {
	router := acp.NewRouter()
	n.acpRouter = router
	n.acpTransport = acp.NewTransport(n.host, n.keyPair.DID, n.keyPair.Private, n, router, n.log)
	n.host.SetStreamHandler(libp2pproto.ID(acp.ProtocolID), n.acpTransport.ServeStream)
}

// ACPRouter يرجع موجّه مهام ACP
func (n *Node) ACPRouter() *acp.Router {
	return n.acpRouter
}

// RegisterACPTask يسجّل معالج مهمة مخصص
func (n *Node) RegisterACPTask(task string, handler acp.TaskHandler) {
	n.acpRouter.Register(task, handler)
}

// SendACPTask يرسل مهمة ACP لنظير
func (n *Node) SendACPTask(ctx context.Context, pid peer.ID, toDID, task string, input interface{}) (*acp.Envelope, error) {
	var raw json.RawMessage
	if input != nil {
		b, err := json.Marshal(input)
		if err != nil {
			return nil, err
		}
		raw = b
	}
	return n.acpTransport.SendTask(ctx, pid, toDID, task, raw, "")
}

// SupportedACPTasks يرجع المهام المدعومة محلياً
func (n *Node) SupportedACPTasks() []string {
	return n.acpRouter.SupportedTasks()
}
