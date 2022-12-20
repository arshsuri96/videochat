

//function going to have access to peers,
//going to be method not just a function, called dispatch keep frame

type Room struct {
	Peers *Peers
	Hub   *chat.Hub
}

// connections is for multiple connections such as video chat and normal chat at the same time. 


type Peers struct {
	ListLock    sync.RWMutex
	Connections []PeerConnectionState
	TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}

type PeerConnectionState struct {
	PeerConnection : *webrtc.PeerConnection
	websocket *ThreadSafeWriter
}

type ThreadSafeWriter struct {
	Conn *websocket.Conn
	Mutex sync.Mutex
}

func(t *ThreadSafeWriter) WriteJSON (v interface()) error{
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	return t.Conn.WriteJSON(v)
}

func (p *Peers) AddTrack(t *webrtc.TrackRemote) *webrtc.TrackLocalStaticRTP{

}

func (p *Peers) RemoveTrack(t *webrtc.TrackLocalStaticRTP){

}

//help us to signal to our peers for signal connection
func (p *Peers) SingalPeerConnection(){

}

func (p *Peers) DispatchKeyFrame() {

}


type WebsocketMessage struct {
	Event string `json:"event"`
	Data string `json:"data"`
}