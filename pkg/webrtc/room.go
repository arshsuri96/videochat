package webrtc

import (
	"log"
	"sync"

	"github.com/gofiber/websocket"
)

func RoomConn(c *websocket.Conn, p *Peers) {
	var config webrtc.Configuation

	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		log.Print(err)
		return
	}

	newPeer := PeerConnectionsState{
		PeerConnection: peerConnection,
		WebSocket:      &ThreadSafeWriter{},
		Conn:           c,
		Mutex:          sync.Mutex{},
	}
}
