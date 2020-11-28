package main

type Hub struct {
	rooms      map[*Room]bool
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

func newHub() *Hub {
	return &Hub{
		rooms:      make(map[*Room]bool),
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}

func (hub *Hub) Run() {
	for {
		select {

		case client := <-hub.register:
			hub.registerClient(client)

		case client := <-hub.unregister:
			hub.unregisterClient(client)

		case message := <-hub.broadcast:
			hub.brocastToClient(message)
		}
	}
}

func (hub *Hub) registerClient(client *Client) {
	hub.clients[client] = true
}

func (hub *Hub) unregisterClient(client *Client) {
	if _, ok := hub.clients[client]; ok {
		delete(hub.clients, client)
	}
}

func (hub *Hub) brocastToClient(message []byte) {
	for client := range hub.clients {
		client.send <- message
	}
}

func (hub *Hub) findRoomByName(name string) *Room {
	var foundRoom *Room
	for room := range hub.rooms {
		if room.GetName() == name {
			foundRoom = room
			break
		}
	}

	return foundRoom
}

func (hub *Hub) createRoom(name string) *Room {
	room := NewRoom(name)
	go room.RunRoom()
	hub.rooms[room] = true

	return room
}
