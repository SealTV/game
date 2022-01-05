package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sealtv/game/game"
	"github.com/sealtv/game/types"
)

const (
	atlasPath = "./assets/sprites.xml"
)

func main() {
	atlas, err := types.LoadTextureAtlas(atlasPath)
	if err != nil {
		log.Fatal(err)
	}

	game := game.NewGame(atlas)
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(1048, 728)
	ebiten.SetWindowTitle("Math Game")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
