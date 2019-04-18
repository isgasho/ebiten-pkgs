package main

import (
	"image/color"

	im "github.com/Rosalita/my-ebiten/pkgs/imagemenu"
	"github.com/Rosalita/my-ebiten/resources/ui"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil" // This is required to draw debug texts.
	"github.com/hajimehoshi/ebiten/inpututil"  // required for isKeyJustPressed
)

var (
	imgMenu im.ImageMenu
)

func update(screen *ebiten.Image) error {

	screen.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff})

	ebitenutil.DebugPrint(screen, "Image Menu - use left and right arrow keys")

	imgMenu.Draw(screen)

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		imgMenu.IncrementSelected()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		imgMenu.DecrementSelected()
	}

	return nil
}

func main() {

	menuItems := []im.Item{
		{
			Name:  "human",
			Bytes: ui.Human_s,
		},
		{
			Name:  "creature",
			Bytes: ui.Creature_s,
		},
	}

	menuInput := im.Input{
		Tx:        100,
		Ty:        36,
		ImgWidth:  100,
		ImgHeight: 100,
		Items:     menuItems,
	}

	imgMenu, _ = im.NewMenu(menuInput)

	if err := ebiten.Run(update, 320, 240, 2, "Image menu"); err != nil {
		panic(err)
	}
}
