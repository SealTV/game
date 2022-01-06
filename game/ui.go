package game

import (
	"fmt"
	"image/color"

	"github.com/blizzy78/ebitenui"
	"github.com/blizzy78/ebitenui/image"
	"github.com/blizzy78/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"github.com/sealtv/game/types"
	"golang.org/x/image/font"
	fonts "golang.org/x/mobile/exp/font"
)

type gameUI struct {
	ui      *ebitenui.UI
	buttons [4]*widget.Button
}

func newUI(atlas types.TextureAtlas) (*ebitenui.UI, error) {

	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Stretch([]bool{true}, []bool{false, true, false}),
			widget.GridLayoutOpts.Padding(widget.Insets{
				Top:    0,
				Bottom: 0,
			}),
			widget.GridLayoutOpts.Spacing(0, 24))),
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

func createButtons(c *widget.Container, atlas types.TextureAtlas) [4]*widget.Button {
	result := [4]*widget.Button{}

	img := image.NewNineSlice(
		atlas.Textures[3].Img,
		[3]int{6, 11, 6},
		[3]int{6, 20, 21},
	)

	ttfBytes := fonts.Default
	t, _ := truetype.Parse(ttfBytes())
	f := truetype.NewFace(t, &truetype.Options{
		Size:    14,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	// f := font.Drawer
	for i := 0; i < 4; i++ {
		clickHandler := func(i int) func(args *widget.ButtonClickedEventArgs) {
			return func(args *widget.ButtonClickedEventArgs) {
				fmt.Println("click: ", i)
			}
		}(i)
		b := widget.NewButton(
			widget.ButtonOpts.ClickedHandler(clickHandler),
			// widget.ButtonOpts.GraphicNineSlice(img),
			widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(
				widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
					Stretch:   false,
					Position:  widget.RowLayoutPositionStart,
					MaxWidth:  100,
					MaxHeight: 100,
				})),
			)),
			widget.ButtonOpts.Text(fmt.Sprintf("Button %d", i+1), f, &widget.ButtonTextColor{
				Idle:     color.Black,
				Disabled: color.Opaque,
			}),
			widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(10)),
			widget.ButtonOpts.Image(&widget.ButtonImage{
				Idle:     img,
				Hover:    img,
				Pressed:  img,
				Disabled: img,
			}),
		)

		c.AddChild(b)
		result[i] = b
	}

	return result
}

func getHeaderWidget(atlas types.TextureAtlas) widget.PreferredSizeLocateableWidget {
	img := image.NewNineSlice(
		atlas.Textures[len(atlas.Textures)-1].Img,
		[3]int{82, 460, 82},
		[3]int{77, 0, 0},
	)

	c := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top:    5,
				Left:   10,
				Right:  10,
				Bottom: 5,
			}),
			widget.RowLayoutOpts.Spacing(10),
		)),
		widget.ContainerOpts.BackgroundImage(img),
	)
	createButtons(c, atlas)

	return c
}

func getContent(atlas types.TextureAtlas) widget.PreferredSizeLocateableWidget {
	img := image.NewNineSlice(
		atlas.Textures[len(atlas.Textures)-3].Img,
		[3]int{82, 460, 82},
		[3]int{45, 201, 45},
	)

	c := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
				widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(0)),
				widget.RowLayoutOpts.Spacing(0),
			),
		),
		widget.ContainerOpts.BackgroundImage(img),
	)
	return c
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
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{
					HorizontalPosition: widget.AnchorLayoutPositionCenter,
					VerticalPosition:   widget.AnchorLayoutPositionEnd,
					StretchHorizontal:  true,
					StretchVertical:    false,
				},
			),
		),
	)
}
