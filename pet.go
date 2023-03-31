package main

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	_ "github.com/polpettone/desktop-pet/statik"
	"github.com/rakyll/statik/fs"
)

type Pet struct {
	X                int
	Y                int
	Images           []*ebiten.Image
	CurrentImage     *ebiten.Image
	AnimationCounter int
	Scale            float64
}

func NewPet(x, y int, scale float64, path string) *Pet {

	images, err := loadImagesFromStatik()

	if err != nil {
		log.Fatal(err)
	}

	return &Pet{
		X:                x,
		Y:                y,
		Images:           images,
		AnimationCounter: 0,
		CurrentImage:     images[0],
		Scale:            scale,
	}
}

func (p *Pet) Draw(screen *ebiten.Image, gameClock int) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(p.Scale, p.Scale)
	op.GeoM.Translate(float64(p.X), float64(p.Y))

	screen.DrawImage(p.CurrentImage, op)

	if gameClock%40 == 0 {
		p.CurrentImage = p.Images[p.AnimationCounter]
		p.AnimationCounter = p.AnimationCounter + 1
		if p.AnimationCounter >= len(p.Images) {
			p.AnimationCounter = 0
		}
	}

}

func loadImagesFromStatik() ([]*ebiten.Image, error) {
	images := []*ebiten.Image{}

	files := []string{"hamster-1.png", "hamster-2.png", "hamster-2.png", "hamster-4.png"}

	statikFS, err := fs.New()
	if err != nil {
		return nil, err
	}

	path := "pet-2"

	for _, file := range files {

		filePath := "/" + path + "/" + file
		log.Printf("open %s", filePath)
		r, err := statikFS.Open(filePath)
		if err != nil {
			return nil, err
		}
		defer r.Close()

		img, _, err := image.Decode(r)
		img2, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)

		if err != nil {
			return nil, err
		}
		images = append(images, img2)
	}

	return images, nil
}

func loadImages(path string) ([]*ebiten.Image, error) {
	images := []*ebiten.Image{}

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		img, _, err := ebitenutil.NewImageFromFile(path+"/"+file.Name(), ebiten.FilterDefault)
		if err != nil {
			return nil, err
		}
		images = append(images, img)
	}

	return images, nil
}
