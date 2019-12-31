package actions

import (
	"kiv_ups_server/masterserver/interfaces"
	"math/rand"
)

func ConvertShadowPlayerToPlayer(player interfaces.Player, name string) interfaces.Player {
	player.SetUID(interfaces.PlayerUID(rand.Int()))
	player.SetName(name)

	return player
}