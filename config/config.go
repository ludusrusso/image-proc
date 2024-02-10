package config

import (
	"fmt"
	"strings"
)

type CropType string

const (
	CropTypeScale CropType = "scale"
	CropTypeCrop  CropType = "crop"
	CropTypeThumb CropType = "thumb"
)

type GravityType string

const (
	GravityTypeFace   GravityType = "face"
	GravityTypeCenter GravityType = "center"
)

type ProcConfig struct {
	Crop    CropType
	Gravity GravityType
	Width   uint
	Height  uint
}

func (c ProcConfig) String() string {
	return fmt.Sprintf("c_%s,g_%s,w_%d,h_%d", c.Crop, c.Gravity, c.Width, c.Height)
}

func Parse(cnfs string) ProcConfig {
	parsers := map[string]paraseConfg{
		"crop":   parseCropConfig,
		"height": parseHeightConfig,
		"width":  parseWidthConfig,
		"grav":   parseGravityConfig,
	}

	c := ProcConfig{
		Crop:    CropTypeScale,
		Gravity: GravityTypeCenter,
	}
	for _, cnf := range strings.Split(cnfs, ",") {
		for _, parser := range parsers {
			ok := parser(cnf, &c)
			if ok {
				break
			}
		}
	}
	return c
}
