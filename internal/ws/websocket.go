package ws

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// HandleConversion handles incoming WebSocket connections for audio conversion
func HandleConversion(c *fiber.Ctx) error {
	conn, err := upgrader.Upgrade(c.Context(), nil)
	if err != nil {
		return fmt.Errorf("could not upgrade connection: %w", err)
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return err
		}

		flacData, err := ConvertWAVToFLAC(msg)
		if err != nil {
			return err
		}

		if err := conn.WriteMessage(websocket.BinaryMessage, flacData); err != nil {
			return err
		}
	}
}

// ConvertWAVToFLAC converts WAV data to FLAC (dummy implementation)
func ConvertWAVToFLAC(wavData []byte) ([]byte, error) {
	// Here you would implement the actual conversion logic using a suitable library
	return wavData, nil // This should return the actual converted FLAC data
}
