package maps

import (
	"errors"
)

type GroundType string

const (
	Snow         GroundType = "Snowtl"
	Grass        GroundType = "GRASTL"
	Sand         GroundType = "sandtl"
	Swamp        GroundType = "Swmptl"
	Dirt         GroundType = "DIRTTL"
	Water        GroundType = "Watrtl"
	Subterranean GroundType = "Subbtl"
)

type Ground struct {
	Type   GroundType
	Frames []int
}

var ErrUndefinedGroundType = errors.New("undefined map ground type")

func parseGround(name string) (Ground, error) {
	gt, err := getGroundTypeFromMapMetaName(name)
	if err != nil {
		return Ground{}, err
	}

	f := getFramesForGround(gt)

	return Ground{
		Type:   gt,
		Frames: f,
	}, nil
}

func getFramesForGround(gt GroundType) []int {
	switch gt {
	case Grass:
		return []int{0, 1, 2, 3, 4, 5, 6, 7, 10, 11, 12, 13}
	default:
		return []int{0}
	}
}

func getGroundTypeFromMapMetaName(name string) (GroundType, error) {
	switch name {
	case "SNOW":
		return Snow, nil
	case "GRASS":
		return Grass, nil
	case "SAND":
		return Sand, nil
	case "SWAMP":
		return Swamp, nil
	case "DIRT":
		return Dirt, nil
	case "SUBTERRANEAN":
		return Subterranean, nil
	case "WATER":
		return Water, nil
	default:
		return "", ErrUndefinedGroundType
	}
}
