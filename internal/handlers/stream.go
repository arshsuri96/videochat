package handlers

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber"
	"github.com/gofiber/websocket/v2"
)

func Stream(c *fiber.Ctx) error {
	suuid := c.Params("suuid")
	if suuid == "" {
		c.Status(400)
		return nil
	}
	ws := "ws"
	if os.Getenv("ENVIRONMENT") == "PRODUCTION" {
		ws = "wss"
	}

	w.RoomsLock.Lock()
	if _, ok := w.Stream[suuid]; !ok {
		w.RoomsLock.Unlock()
		return c.Render("stream", fiber.Map{
			"StreamWebsocketAddress": fmt.Sprintf("%s://%s/stream/%s/websocket", ws, c.Hostname(), suuid),
			"ChatWebsocketAddr":      fmt.Sprintf("%s://%s/stream/%s/chat/websocket", ws, c.Hostanme(), suuid),
			"ViewerWebsocketAddr":    fmt.Sprintf("%s://%s/stream/%s/chat/websocket", ws, c.Hostname(), suuid),
			"Type":                   "stream",
		}, "layouts/main")
	}
	w.RoomsLock.Unlock()
	return c.Render("stream", fiber.Map{
		"NoStream": "true",
		"Leave":    "true",
	}, "layouts/main")
}

func StreamWebsocket(c *websocket.Conn) {

}

func StreamViewerWebsocket(c *websocket.Conn) {

}

func ViewerConn(c *websocket.Conn, p *w.Peers) {

}
