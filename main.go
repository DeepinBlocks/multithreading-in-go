package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
	"multithreading-in-go/boids"
)

var (
	green = color.RGBA{R: 10, G: 255, B: 50, A: 255}
)

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, boid := range boids.Boids {
		screen.Set(int(boid.Position.X+1), int(boid.Position.Y), green)
		screen.Set(int(boid.Position.X-1), int(boid.Position.Y), green)
		screen.Set(int(boid.Position.X), int(boid.Position.Y-1), green)
		screen.Set(int(boid.Position.X), int(boid.Position.Y+1), green)
	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return boids.ScreenWidth, boids.ScreenHeight
}

func main() {
	for i, row := range boids.BoidMap {
		for j := range row {
			boids.BoidMap[i][j] = -1
		}
	}

	for i := 0; i < boids.BoidCount; i++ {
		boids.CreateBoid(i)
	}

	ebiten.SetWindowSize(boids.ScreenWidth*2, boids.ScreenHeight*2)
	ebiten.SetWindowTitle("Boids in a box")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatalln(err)
	}
}
