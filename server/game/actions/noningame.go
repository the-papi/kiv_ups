package actions

import (
	"kiv_ups_server/game/interfaces"
	"kiv_ups_server/net/tcp"
	"kiv_ups_server/net/tcp/protocol"
)

func (a KeepAliveAction) Process(s interfaces.MasterServer, m interfaces.PlayerMessage) ActionResponse {
	keepAliveData := m.GetMessage().Message.(*protocol.KeepAliveMessage)

	if keepAliveData.Ping == "pong" {
		return ActionResponse{
			ServerMessage: tcp.ServerMessage{
				Data:        &protocol.KeepAliveMessage{Ping:"ping-pong"},
				Status:      true,
				Message:     "",
			},
			Targets: []interfaces.Player{m.GetPlayer()},
		}
	} else {
		return ActionResponse{
			ServerMessage: tcp.ServerMessage{
				Data:        nil,
				Status:      false,
				Message:     "",
			},
			Targets: []interfaces.Player{m.GetPlayer()},
		}
	}
}

func (a AuthenticateAction) Process(s interfaces.MasterServer, m interfaces.PlayerMessage) ActionResponse {
	authenticateData := m.GetMessage().Message.(*protocol.AuthenticateMessage)
	s.Authenticate(ConvertShadowPlayerToPlayer(m.GetPlayer(), authenticateData.Name))

	return ActionResponse{
		ServerMessage: tcp.ServerMessage{
			Data:        authenticateData,
			Status:      true,
			Message:     "",
		},
		Targets: []interfaces.Player{m.GetPlayer()},
	}
}