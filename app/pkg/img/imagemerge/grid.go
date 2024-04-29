package imagemerge

import (
	"image"
	"image/color"
	"pasha_bot/pkg/img/imagefilter"
)

// Grid holds the data for each grid
type Grid struct {
	Image           image.Image
	ImageFilePath   string
	BackgroundColor color.Color
	OffsetX         int
	OffsetY         int
	Grids           []*Grid
	Filters         []imagefilter.FilterType
}

// ApplyFilters applies the selected filters to the image.
func (g Grid) ApplyFilters(img image.Image) image.Image {
	for _, filterType := range g.Filters {
		filter := imagefilter.Get(filterType)
		if filter != nil {
			img = filter.Apply(img)
		}
	}

	return img
}
