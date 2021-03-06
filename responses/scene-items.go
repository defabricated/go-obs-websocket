package responses

import (
	"github.com/czlowiekenigma/go-obs-websocket/scene/item"
	"github.com/czlowiekenigma/go-obs-websocket/transforms"
)

type GetSceneItemProperties struct {
	*ResponseBase
	Name     string              `json:"name"`
	Position transforms.Position `json:"position"`
	Rotation float64             `json:"rotation"`
	Scale    transforms.Scale    `json:"scale"`
	Crop     transforms.Crop     `json:"crop"`
	Visible  bool                `json:"visible"`
	Bounds   transforms.Bounds   `json:"bounds"`
}

type DuplicateSceneItem struct {
	*ResponseBase
	Scene string     `json:"scene"`
	Item  item.Item `json:"item"`
}
