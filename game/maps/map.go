package maps

import (
	"fmt"
	"game_try/game/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"gopkg.in/yaml.v3"
	_ "image/png"
)

type GameMap struct {
	Size          int
	Cells         []GameMapCell
	Players       int
	AssetsStore   map[string]*ebiten.Image
	MapStaticView *ebiten.Image
	TickCount     int
}

type GameMapCell struct {
	Ground   Ground
	Resource Resource
}

func (gm *GameMap) DrawMap() (*ebiten.Image, error) {
	img := ebiten.NewImageFromImage(gm.MapStaticView)

	for y := 0; y < gm.Size; y++ {
		for x := 0; x < gm.Size; x++ {
			cellIndex := (y * gm.Size) + x
			cell := gm.Cells[cellIndex]
			err := gm.drawResource(img, cell, x, y)
			if err != nil {
				return nil, err
			}
		}
	}
	ebitenutil.DebugPrint(img, fmt.Sprintf("%f", ebiten.CurrentFPS()))
	return img, nil
}

func (gm *GameMap) drawResource(img *ebiten.Image, cell GameMapCell, x int, y int) error {
	if cell.Resource.Type == "" {
		return nil
	}
	frameIndex := (gm.TickCount / 20) % len(cell.Resource.Frames)
	frame := cell.Resource.Frames[frameIndex]
	assetKey := fmt.Sprintf("%s-%d", cell.Resource.Type, frame)
	frameImg, ok := gm.AssetsStore[assetKey]
	if !ok {
		return fmt.Errorf("no assets found for ground(%s) frame(%d)", cell.Resource.Type, frame)
	}
	dio := ebiten.DrawImageOptions{}
	dio.GeoM.Translate(float64(x*32-32), float64(y*32))
	img.DrawImage(ebiten.NewImageFromImage(frameImg), &dio)
	return nil
}

func GenerateMapFromYaml(yamlMap []byte, assetsLoader assets.Loader) (*GameMap, error) {
	var gameMapMeta struct {
		Size    int `yaml:"size"`
		Players int `yaml:"players"`
		Cells   []struct {
			Ground   string `yaml:"ground"`
			Resource struct {
				Type   string `yaml:"type"`
				Amount int    `yaml:"amount"`
			} `yaml:"resource"`
		} `yaml:"cells"`
	}

	err := yaml.Unmarshal(yamlMap, &gameMapMeta)
	if err != nil {
		return nil, fmt.Errorf("can't parse map yaml: %w", err)
	}
	if gameMapMeta.Size*gameMapMeta.Size != len(gameMapMeta.Cells) {
		return nil, fmt.Errorf("cells count do not match the size of the map, should be %d, got %d", gameMapMeta.Size*2, len(gameMapMeta.Cells))
	}

	mapAssets := map[string]*ebiten.Image{}
	var cells []GameMapCell
	for _, metaCell := range gameMapMeta.Cells {
		g, err := parseGround(metaCell.Ground)
		if err != nil {
			return nil, err
		}
		if g.Type != "" {
			err = assetsLoader.LoadMapGround(string(g.Type), g.Frames, mapAssets)
			if err != nil {
				return nil, fmt.Errorf("can't load ground assets: %s", g.Type)
			}
		}

		r, err := parseResource(metaCell.Resource.Type)
		if err != nil {
			return nil, err
		}
		if r.Type != "" {
			err = assetsLoader.LoadMapResource(string(r.Type), r.Frames, mapAssets)
			if err != nil {
				return nil, fmt.Errorf("can't load resource assets: %s", r.Type)
			}
		}
		r.Amount = metaCell.Resource.Amount

		cells = append(cells, GameMapCell{
			Ground:   g,
			Resource: r,
		})
	}

	mapSurface, err := drawMapStaticView(gameMapMeta.Size, cells, mapAssets)
	if err != nil {
		return nil, fmt.Errorf("can't draw map surface: %w", err)
	}

	return &GameMap{
		Size:          gameMapMeta.Size,
		Players:       gameMapMeta.Players,
		Cells:         cells,
		AssetsStore:   mapAssets,
		MapStaticView: mapSurface,
	}, nil
}

func drawMapStaticView(size int, cells []GameMapCell, assetsStore map[string]*ebiten.Image) (*ebiten.Image, error) {
	img := ebiten.NewImage(size*32, size*32)

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			cellIndex := (y * size) + x
			cell := cells[cellIndex]
			err := drawGround(img, cell, x, y, assetsStore)
			if err != nil {
				return nil, err
			}
		}
	}
	return img, nil
}

func drawGround(img *ebiten.Image, cell GameMapCell, x int, y int, assetsStore map[string]*ebiten.Image) error {
	frame := cell.Ground.Frames[x*y%len(cell.Ground.Frames)]
	frameKey := fmt.Sprintf("%s-%d", cell.Ground.Type, frame)
	frameImg, ok := assetsStore[frameKey]
	if !ok {
		return fmt.Errorf("no assets found for ground(%s) frame(%d)", cell.Ground.Type, frame)
	}
	dio := ebiten.DrawImageOptions{}
	dio.GeoM.Translate(float64(x*32), float64(y*32))
	img.DrawImage(frameImg, &dio)
	return nil
}
