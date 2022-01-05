package game

import (
	"github.com/blizzy78/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sealtv/game/types"
)

type game struct {
	atlas types.TextureAtlas
	ui    *ebitenui.UI
}

func NewGame(atlas types.TextureAtlas) ebiten.Game {
	ui, _ := newUI(atlas)

	return &game{
		atlas: atlas,
		ui:    ui,
	}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *game) Update() error {
	// Write your game's logical update.

	g.ui.Update()

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.

	g.ui.Draw(screen)
	// g.drawMenu(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1048, 728
}
