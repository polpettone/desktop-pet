package main

import (
	"log"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 200
	screenHeight = 200
)

func main() {

	petImagesPath := "pet-2"
	scale := 1.0
	walkMode := Bottom
	if len(os.Args) > 1 {
		petImagesPath = os.Args[1]
	}

	if len(os.Args) > 2 {
		scaleFromArg, err := strconv.ParseFloat(os.Args[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		scale = scaleFromArg
	}

	if len(os.Args) > 3 {
		if os.Args[3] == "RANDOM" {
			walkMode = Random
		}
	}

	pet := NewPet(50, 50, scale, petImagesPath)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Desktop Pet")
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetWindowPosition(1000, 1000)
	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowFloating(true)

	game := &Game{
		Pet:      pet,
		X:        500,
		Y:        500,
		Clock:    0,
		WalkMode: walkMode,
	}
	ebiten.SetScreenTransparent(true)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
