package interfaces

import (
	"kiv_ups_server/net/tcp"
)

type PlayerContext int
type PlayerUID int

type Player interface {
	GetTCPClient() *tcp.Client
	SetTCPClient(*tcp.Client)
	GetUID() PlayerUID
	SetUID(uid PlayerUID)
	GetName() string
	SetName(name string)
	GetContext() PlayerContext
	SetContext(ctx PlayerContext)
	GetConnectedLobby() *Lobby
	SetConnectedLobby(*Lobby)
	GetGameServer() GameServer
	SetGameServer(gs GameServer)
	LeaveGame()
	IsConnected() bool
	RefreshKeepAlive()
	SetLoggedInMenuContext()
	IncrementCheatCounter()
	ResetCheatCounter()
	IncrementInvalidMessageCounter()
	ResetInvalidMessageCounter()
}

type PlayerMessage interface {
	GetMessage() *tcp.ClientMessage
	GetPlayer() Player
}
