package nodes

import (
	"kiv_ups_server/internal/masterserver/gameserver/settings"
	tree2 "kiv_ups_server/internal/masterserver/gameserver/tree"
	interfaces2 "kiv_ups_server/internal/masterserver/interfaces"
	"kiv_ups_server/internal/net/tcp/protocol"
	"math/rand"
)

const Margin = 100
const BorderWidth = 100
const AsteroidsLimit = 50

type AsteroidWrapper struct {
	Node *tree2.Node `json:"-"`
}

func (a *AsteroidWrapper) Init(node *tree2.Node) {
	a.Node = node
}

func (a *AsteroidWrapper) Process(playerMessages []interfaces2.PlayerMessage, delta float64) {
	if len(a.Node.Children) < AsteroidsLimit {
		// create new asteroid
		// randomize pos x
		var minX, maxX, minY, maxY int
		side := rand.Int() % 4
		if side == 0 {
			// up
			minX = 0 - Margin - BorderWidth
			maxX = settings.Width + Margin + BorderWidth
			minY = 0 - Margin - BorderWidth
			maxY = 0 - Margin
		} else if side == 1 {
			// right
			minX = settings.Width + Margin
			maxX = settings.Width + Margin + BorderWidth
			minY = 0 - Margin - BorderWidth
			maxY = settings.Height + Margin + BorderWidth
		} else if side == 2 {
			// bottom
			minX = 0 - Margin - BorderWidth
			maxX = settings.Width + Margin + BorderWidth
			minY = settings.Height + Margin
			maxY = settings.Height + Margin + BorderWidth
		} else if side == 3 {
			// left
			minX = 0 - Margin - BorderWidth
			maxX = 0 - Margin
			minY = 0 - Margin - BorderWidth
			maxY = settings.Width + Margin + BorderWidth
		}

		// Create new asteroid at random position of random scale
		posX := rand.Intn(maxX-minX) + minX
		posY := rand.Intn(maxY-minY) + minY
		node := tree2.NewNode(a.Node, &Asteroid{
			PosX:      float64(posX),
			PosY:      float64(posY),
			VelocityX: -100 + rand.Float64()*(100+100),
			VelocityY: -100 + rand.Float64()*(100+100),
			Rotation:  0,
			Scale:     rand.Float64(),
			Node:      nil,
		})

		node.Init()
		node.Value.Init(&node)

		a.Node.Children = append(a.Node.Children, &node)
	}
}

func (a *AsteroidWrapper) ListenMessages() []protocol.Message {
	return []protocol.Message{}
}

func (a *AsteroidWrapper) Filter(playerMessages []interfaces2.PlayerMessage) []interfaces2.PlayerMessage {
	return playerMessages
}

type Asteroid struct {
	PosX      float64     `json:"pos_x"`
	PosY      float64     `json:"pos_y"`
	VelocityX float64     `json:"velocity_x"`
	VelocityY float64     `json:"velocity_y"`
	Rotation  float64     `json:"rotation"`
	Scale     float64     `json:"scale"`
	Value     int         `json:"value"`
	Radius    float64     `json:"-"`
	Node      *tree2.Node `json:"-"`
}

func (a *Asteroid) Init(node *tree2.Node) {
	a.Node = node
	a.Value = int(100 * a.Scale)
	a.Radius = 100 * a.Scale
}

func (a *Asteroid) Process(playerMessages []interfaces2.PlayerMessage, delta float64) {
	a.PosX += a.VelocityX * delta
	a.PosY += a.VelocityY * delta

	if a.PosX < -Margin-BorderWidth || a.PosX > settings.Width+Margin+BorderWidth ||
		a.PosY < -Margin-BorderWidth || a.PosY > settings.Height+Margin+BorderWidth {
		a.Node.Destroy()
		return
	}
}

func (a *Asteroid) ListenMessages() []protocol.Message {
	return []protocol.Message{}
}

func (a *Asteroid) Filter(playerMessages []interfaces2.PlayerMessage) []interfaces2.PlayerMessage {
	return playerMessages
}
