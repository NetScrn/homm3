package assets

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

type Loader interface {
	LoadMapGround(name string, frames []int, assetsStore map[string]*ebiten.Image) error
	LoadMapResource(name string, frames []int, assetsStore map[string]*ebiten.Image) error
}

type DriveLoader struct {
}

func (h DriveLoader) LoadMapGround(name string, frames []int, assetsStore map[string]*ebiten.Image) error {
	return loadFrames(name, "grounds_map", frames, assetsStore)
}

func (h DriveLoader) LoadMapResource(name string, frames []int, assetsStore map[string]*ebiten.Image) error {
	return loadFrames(name, "resources_map", frames, assetsStore)
}

func loadFrames(name string, resDir string, frames []int, assetsStore map[string]*ebiten.Image) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	framesPath := filepath.Join(dir, "assets", resDir, name)
	for _, frame := range frames {
		framePath := filepath.Join(framesPath, strconv.Itoa(frame)+".png")
		b, err := ioutil.ReadFile(framePath)
		if err != nil {
			return err
		}
		img, _, err := image.Decode(bytes.NewReader(b))
		eImg := ebiten.NewImageFromImage(img)

		storeKey := fmt.Sprintf("%s-%d", name, frame)
		if _, ok := assetsStore[storeKey]; !ok {
			assetsStore[storeKey] = eImg
		}
	}
	return nil
}
