package interfaces

import (
	"kiv_ups_server/net/tcp"
)

type Lobby struct {
	Name         string
	Owner        Player
	Players      map[PlayerUID]Player
	PlayersLimit int
}

func (l *Lobby) KickPlayer(player Player){
	panic("implement me")
}

func (l *Lobby) KickPlayers() {
	panic("implement me")
}

type MasterServer interface {
	Start() (err error)
	Stop() (err error)
	RunAction(message tcp.ClientMessage) (err error)
	GetTCPServer() *tcp.Server
	GetPlayers() map[tcp.UID]Player
	Authenticate(player Player)
	AddLobby(lobby *Lobby)
	DeleteLobby(name string)
	GetLobby(name string) (*Lobby, error)
	GetLobbies() []*Lobby
}
