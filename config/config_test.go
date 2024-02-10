package config_test

import (
	"testing"

	"github.com/ludusrusso/image-proc/config"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	cases := []struct {
		in   string
		want config.ProcConfig
	}{
		{"c_crop", config.ProcConfig{Crop: config.CropTypeCrop, Gravity: config.GravityTypeCenter, Width: 0, Height: 0}},
		{"c_scale", config.ProcConfig{Crop: config.CropTypeScale, Gravity: config.GravityTypeCenter, Width: 0, Height: 0}},
		{"c_crop,w_100,h_100", config.ProcConfig{Crop: config.CropTypeCrop, Gravity: config.GravityTypeCenter, Width: 100, Height: 100}},
		{"c_scale,w_100,h_100", config.ProcConfig{Crop: config.CropTypeScale, Gravity: config.GravityTypeCenter, Width: 100, Height: 100}},
		{"g_center,c_crop,w_100,h_100", config.ProcConfig{Crop: config.CropTypeCrop, Width: 100, Height: 100, Gravity: config.GravityTypeCenter}},
		{"", config.ProcConfig{Crop: config.CropTypeScale, Gravity: config.GravityTypeCenter, Width: 0, Height: 0}},
		{"g_face", config.ProcConfig{Crop: config.CropTypeScale, Gravity: config.GravityTypeFace, Width: 0, Height: 0}},
		{"g_xxx,c_xxx", config.ProcConfig{Crop: config.CropTypeScale, Gravity: config.GravityTypeCenter, Width: 0, Height: 0}},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			got := config.Parse(c.in)
			assert.Equal(t, c.want, got)
		})
	}
}
