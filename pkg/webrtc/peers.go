package webrtc

import (
	"sync"
	"videochat/pkg/chat"
)

//function going to have access to peers,
//going to be method not just a function, called dispatch keep frame

type Room struct {
	Peers *Peers
	Hub   *chat.Hub
}

type Peers struct {
	ListLock    sync.RWMutex
	Connections []PeerConnectionsState
	TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}

func (p *Peers) DispatchKeyFrame() {

}
