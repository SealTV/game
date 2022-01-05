package game

import (
	"image/color"

	"github.com/blizzy78/ebitenui"
	"github.com/blizzy78/ebitenui/image"
	"github.com/blizzy78/ebitenui/widget"
	"github.com/sealtv/game/types"
)

func newUI(atlas types.TextureAtlas) (*ebitenui.UI, error) {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				StretchHorizontal: true,
				StretchVertical:   true,
			}),
		),
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
				widget.RowLayoutOpts.Spacing(30),
				// widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(40)),
			),
		),
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{R: 240, G: 120, B: 0, A: 128})),
	)

	_ = rootContainer.AddChild(getHeaderWidget(atlas))
	_ = rootContainer.AddChild(getContent(atlas))
	_ = rootContainer.AddChild(getBottomWidget(atlas))

	ui := &ebitenui.UI{
		Container: rootContainer,
	}

	return ui, nil
}

func getHeaderWidget(atlas types.TextureAtlas) widget.PreferredSizeLocateableWidget {
	img := image.NewNineSlice(
		atlas.Textures[len(atlas.Textures)-1].Img,
		[3]int{82, 460, 82},
		[3]int{77, 0, 0},
	)

	return widget.NewGraphic(
		widget.GraphicOpts.ImageNineSlice(img),
		widget.GraphicOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(
				widget.RowLayoutData{
					Stretch: true,
				},
			),
		),
	)
}
func getContent(atlas types.TextureAtlas) widget.PreferredSizeLocateableWidget {
	img := image.NewNineSlice(
		atlas.Textures[len(atlas.Textures)-3].Img,
		[3]int{82, 460, 82},
		[3]int{50, 50, 37},
	)

	return widget.NewGraphic(
		widget.GraphicOpts.ImageNineSlice(img),
		widget.GraphicOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.WidgetOpts.LayoutData(
				widget.RowLayoutData{
					Stretch: true,
				},
			)),
		),
	)
}

func getBottomWidget(atlas types.TextureAtlas) widget.PreferredSizeLocateableWidget {
	img := image.NewNineSlice(
		atlas.Textures[len(atlas.Textures)-2].Img,
		[3]int{82, 460, 82},
		[3]int{0, 0, 77},
	)

	return widget.NewGraphic(
		widget.GraphicOpts.ImageNineSlice(img),
		widget.GraphicOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.WidgetOpts.LayoutData(
				widget.RowLayoutData{
					Stretch: true,
				},
			)),
		),
	)
}
