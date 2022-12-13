package chat

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client

	// concpets around websocket
	// these are channels which will tell if the client is
	// on the page or not, registers - when somebody registers on the room.
}
