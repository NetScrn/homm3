package game

import (
	"fmt"
	"game_try/game/assets"
	"game_try/game/maps"
	"github.com/hajimehoshi/ebiten/v2"
)

type Homm3Game struct {
	gameMap      *maps.GameMap
	translations map[string]string
	tickCount    int
}

func (g *Homm3Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 32 * 20, 32 * 20
}

func (g *Homm3Game) Update() error {
	if g.tickCount == 0 {
		g.tickCount = 1
	} else {
		g.tickCount++
	}
	g.gameMap.TickCount = g.tickCount
	return nil
}

func (g *Homm3Game) Draw(screen *ebiten.Image) {
	mapImg, err := g.gameMap.DrawMap()
	if err != nil {
		panic(err)
	}

	dio := &ebiten.DrawImageOptions{}
	screen.DrawImage(mapImg, dio)
}

func NewHoMM3Game(mapYaml []byte, translations map[string]string) (*Homm3Game, error) {
	gameMap, err := maps.GenerateMapFromYaml(mapYaml, assets.DriveLoader{})
	if err != nil {
		return nil, fmt.Errorf("can't parses game map: %w", err)
	}

	return &Homm3Game{
		gameMap:      gameMap,
		translations: translations,
	}, nil
}
