package creatures

import "image"

type GameFrames []image.Image

type CreatureStack struct {
	Name       string
	MapFrames  GameFrames
	MoveFrames GameFrames
	IdleFrames GameFrames
}

func NewCreatureStack(cmf creatureMetaFile, count int) {

}
