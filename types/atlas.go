package types

import (
	"encoding/xml"
	"fmt"
	"image"
	_ "image/png"
	"io"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
)

type TextureAtlas struct {
	XMLName   xml.Name  `xml:"TextureAtlas"`
	Textures  []Texture `xml:"SubTexture"`
	ImagePath string    `xml:"imagePath,attr"`
}

type Texture struct {
	XMLName xml.Name      `xml:"SubTexture"`
	Name    string        `xml:"name,attr"`
	X       int           `xml:"x,attr"`
	Y       int           `xml:"y,attr"`
	Width   int           `xml:"width,attr"`
	Height  int           `xml:"height,attr"`
	Img     *ebiten.Image `xml:"-"`
}

func LoadTextureAtlas(file string) (TextureAtlas, error) {
	xF, err := os.Open(file)
	if err != nil {
		return TextureAtlas{}, fmt.Errorf("can't load atlas file")
	}
	defer xF.Close()

	atlas, err := parseTextureAtlas(xF)
	if err != nil {
		return TextureAtlas{}, fmt.Errorf("can't parse texture atlas")
	}

	sF, err := os.Open(fmt.Sprintf("%s/%s", filepath.Dir(file), atlas.ImagePath))
	if err != nil {
		return TextureAtlas{}, fmt.Errorf("can't open atlas image: %w", err)
	}
	defer sF.Close()

	img, err := loadImage(sF)
	if err != nil {
		return TextureAtlas{}, fmt.Errorf("can't load atlas image: %w", err)
	}

	eImage := ebiten.NewImageFromImage(img)

	for i := range atlas.Textures {
		t := atlas.Textures[i]

		img := (eImage.SubImage(image.Rectangle{
			Min: image.Point{X: t.X, Y: t.Y},
			Max: image.Point{X: t.X + t.Width, Y: t.Y + t.Height - 1},
		})).(*ebiten.Image)

		atlas.Textures[i].Img = img
	}

	return atlas, nil
}

func parseTextureAtlas(in io.Reader) (TextureAtlas, error) {
	result := TextureAtlas{}
	decoder := xml.NewDecoder(in)
	if err := decoder.Decode(&result); err != nil {
		return TextureAtlas{}, fmt.Errorf("can't decode xml data: %w", err)
	}

	return result, nil
}

func loadImage(r io.Reader) (image.Image, error) {
	i, _, err := image.Decode(r)
	if err != nil {
		return nil, fmt.Errorf("can't decode image: %w", err)
	}

	return i, nil
}
