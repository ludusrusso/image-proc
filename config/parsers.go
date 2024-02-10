package config

import (
	"strconv"
	"strings"
)

type paraseConfg = func(cnf string, c *ProcConfig) bool

func parseCropConfig(cnf string, c *ProcConfig) bool {
	crop, ok := checkConfig(cnf, "c_")
	if !ok {
		return false
	}

	cropTypes := []CropType{
		CropTypeScale,
		CropTypeCrop,
		CropTypeThumb,
	}

	for _, t := range cropTypes {
		if crop == string(t) {
			c.Crop = t
			return true
		}
	}

	return true
}

func parseGravityConfig(cnf string, c *ProcConfig) bool {
	grav, ok := checkConfig(cnf, "g_")
	if !ok {
		return false
	}

	gTypes := []GravityType{
		GravityTypeFace,
		GravityTypeCenter,
	}

	for _, t := range gTypes {
		if grav == string(t) {
			c.Gravity = t
			return true
		}
	}

	return true
}

func parseWidthConfig(cnf string, c *ProcConfig) bool {
	ws, ok := checkConfig(cnf, "w_")
	if !ok {
		return false
	}

	w, err := strconv.Atoi(ws)
	if err == nil {
		c.Width = uint(w)
	}

	return true
}

func parseHeightConfig(cnf string, c *ProcConfig) bool {
	hs, ok := checkConfig(cnf, "h_")
	if !ok {
		return false
	}

	w, err := strconv.Atoi(hs)
	if err == nil {
		c.Height = uint(w)
	}

	return true
}

func checkConfig(cnf string, prefix string) (string, bool) {
	if !strings.HasPrefix(cnf, prefix) {
		return "", false
	}

	return strings.Replace(cnf, prefix, "", 1), true
}
