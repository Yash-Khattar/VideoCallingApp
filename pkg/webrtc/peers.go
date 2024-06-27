type Room struct{
	Peers *Peers
	Hub *chat.Hub
}

type Peers struct {
	ListLock sync.PWMutex
	Connections []PeerConnectionState
	TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}

func (p *Peers) DispatchKeyFrames(){

}