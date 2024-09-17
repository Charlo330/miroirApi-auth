// Package controllers provides functions to handle desktop app authentication.
// It includes functions to handle WebSocket connections, generate unique IDs, and handle desktop app login requests.

package controllers

import (
	"miroirapiauth/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// upgrader is a WebSocket upgrader with read and write buffer sizes of 1024.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections by default.
		return true
	},
}

// wsList is a map of WebSocket connections with unique IDs as keys.
var wsList = map[int]*websocket.Conn{}

// getLoginId generates a unique ID and adds the WebSocket connection to the wsList map.
func GetLoginId(c *gin.Context) {
	// Upgrade the HTTP connection to a WebSocket connection.
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
	}
	defer conn.Close()

	// Generate a unique 6-digit ID.
	strId, err := services.GenerateId()

	// Convert the ID to an int.
	id, err := strconv.Atoi(strId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
	}

	// Add the WebSocket connection to the wsList map.
	wsList[id] = conn

	// Set the close handler to remove the connection from the wsList map when the connection is closed.
	conn.SetCloseHandler(func(code int, text string) error {
		delete(wsList, id)
		return nil
	})

	// send the ID to the client.
	conn.WriteMessage(websocket.TextMessage, []byte("id: "+strId))

	// for loop to keep the connection open.
	for {
		_, message, err := conn.ReadMessage()

		if string(message) == "/generateNewId" {
			strId, err := services.GenerateId()
			delete(wsList, id)

			id, err := strconv.Atoi(strId)

			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte("Error"))
			}
			wsList[id] = conn

			// send the ID to the client.
			conn.WriteMessage(websocket.TextMessage, []byte("id: "+strId))
		}

		if err != nil {
			delete(wsList, id)
			return
		}
	}
}

// DesktopLogin sends a token to the WebSocket connection with the specified ID.
func DesktopLogin(c *gin.Context) {
	// Get the ID from the request body.
	var body struct {
		Id string `validate:"required" json:"id"`
	}
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	// Get the token from the request header.
	tokenString := string(c.Request.Header.Get("Authorization")[7:])

	id, err := strconv.Atoi(body.Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id",
		})
		return
	}

	// check in idList if the id is not expired (after 15 minutes)
	connection := wsList[id]

	if !services.CheckifIdIsExpired(id) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "The id is expired",
		})

		connection.WriteMessage(websocket.TextMessage, []byte("id expired"))

		return
	}

	if connection == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id",
		})
		return
	}

	// get the id of the client with his token
	userId, err := services.GetUserIdByToken(tokenString)

	userIdStr := strconv.Itoa(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid token",
		})
		return
	}

	// remove the id from the wsList map.
	delete(wsList, id)
	connection.WriteMessage(websocket.TextMessage, []byte("token: "+tokenString+"\nuser_id: "+userIdStr))
}
