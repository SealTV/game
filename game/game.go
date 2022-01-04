package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sealtv/game/types"
)

type game struct {
	atlas types.TextureAtlas
}

func NewGame(atlas types.TextureAtlas) ebiten.Game {
	return &game{
		atlas: atlas,
	}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *game) Update() error {
	// Write your game's logical update.
	return nil
}

var i int

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.

	g.drawMenu(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func (g *game) drawMenu(screen *ebiten.Image) {
	screen.DrawImage(g.atlas.Textures[len(g.atlas.Textures)-5].Img, &ebiten.DrawImageOptions{})
	i++
	if i >= len(g.atlas.Textures) {
		i = 0
	}

	// header
	screen.DrawImage(g.atlas.Textures[42].Img, nil)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(g.atlas.Textures[42].Width), 0)

	opts.GeoM.Scale(float64(10), 1)

	screen.DrawImage(g.atlas.Textures[43].Img, opts)

	opts = &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(screen.Bounds().Dx()-g.atlas.Textures[44].Width-1), 0)
	screen.DrawImage(g.atlas.Textures[44].Img, opts)
}
