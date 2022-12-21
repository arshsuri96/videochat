

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
	p.ListLock.Lock()
	defer func(){
		p.ListLock.Unlock()
		p.SingalPeerConnections()
	}()

	trackLocal, err := webrtc.NewTrackLocalStaticRTP(t.Codec().RTPCodecCapability, t.ID(). t.StreamID())

	if err!= nil{
		log.Println(err.Error())
		return nil
	}

	p.TrackLocals[t.ID()] = trackLocal
	return trackLocal
}

func (p *Peers) RemoveTrack(t *webrtc.TrackLocalStaticRTP){
	p.ListLock.Lock()
	defer func(){
		p.ListLock.Unlock()
		p.SingalPeerConnections()
	}()
	delete(p.TrackLocals, t.ID[])
}

//help us to signal to our peers for signal connection
func (p *Peers) SingalPeerConnections(){
	p.ListLock.Lock()
	defer func(){
		p.ListLock.Unlock()
		p.DispatchKeyFrame()
	}()

	attemptSync := func (tryAgain bool){
		for i := range p.Connections{
			if p.Connections[i].PeerConnection.PeerConnectionState() == webrtc.PeerConnectionStateClosed{
				p.Connections = append(p.Connections[:i], p.Connections[i+1]...)
				log.Println("a".p.Connections)
				return true
			}

			existingSender := map[string]bool{}
			for _, sender := range p.Connections[i].PeerConnection.GetSender(){
				if sender.Track() == nil{
					continue
				}

				existingSenders[senders.Track().ID()] = true

				if _,ok := p.TrackLocals[sender.Track().ID()]; !ok{
					if err:= p.Connections[i].PeerConnection.RemoveTrack(Sender); err!= nil{
						return true
					}
				}
			}

			
		}
	}
}

func (p *Peers) DispatchKeyFrame() {

}


type WebsocketMessage struct {
	Event string `json:"event"`
	Data string `json:"data"`
}