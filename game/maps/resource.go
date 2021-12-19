package maps

import (
	"errors"
)

type ResourceType string

const (
	Gold     ResourceType = "AVTgold0"
	Wood     ResourceType = "AVTwood0"
	Crystals ResourceType = "AVTCRYS0"
	Gems     ResourceType = "AVTgems0"
	Mercury  ResourceType = "AVTmerc0"
	Ore      ResourceType = "AVTORE0"
	Sulfur   ResourceType = "AVTsulf0"
)

type Resource struct {
	Type   ResourceType
	Frames []int
	Amount int
}

var ErrUndefinedResourceType = errors.New("undefined map resource type")

func parseResource(name string) (Resource, error) {
	rt, err := getResourceTypeFromMapMetaName(name)
	if err != nil {
		return Resource{}, nil
	}

	f := getFramesForResource(rt)

	return Resource{
		Type:   rt,
		Frames: f,
		Amount: 0,
	}, nil

}

func getFramesForResource(rt ResourceType) []int {
	switch rt {
	case Gold:
		return []int{0, 1, 2, 3, 4, 5, 6, 7}
	default:
		return []int{0}
	}
}

func getResourceTypeFromMapMetaName(name string) (ResourceType, error) {
	switch name {
	case "GOLD":
		return Gold, nil
	case "WOOD":
		return Wood, nil
	case "CRYSTALS":
		return Crystals, nil
	case "GEMS":
		return Gems, nil
	case "MERCURY":
		return Mercury, nil
	case "ORE":
		return Ore, nil
	case "SULFUR":
		return Sulfur, nil
	default:
		return "", ErrUndefinedResourceType
	}
}
