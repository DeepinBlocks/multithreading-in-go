package boids

import (
	"math/rand"
	"time"
)

const (
	ScreenWidth, ScreenHeight = 640, 360
	BoidCount                 = 500
)

var (
	Boids [BoidCount]*Boid
)

type Boid struct {
	Position Vector2D
	Velocity Vector2D
	Id       int
}

func (b *Boid) moveOne() {
	b.Position = b.Position.Add(b.Velocity)
	next := b.Position.Add(b.Velocity)
	if next.X >= ScreenWidth || next.X < 0 {
		b.Velocity = Vector2D{-b.Velocity.X, b.Velocity.Y}
	}
	if next.Y >= ScreenHeight || next.Y < 0 {
		b.Velocity = Vector2D{b.Velocity.X, -b.Velocity.Y}
	}
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func CreateBoid(bid int) {
	b := Boid{
		Position: Vector2D{rand.Float64() * ScreenWidth, rand.Float64() * ScreenHeight},
		Velocity: Vector2D{(rand.Float64() * 2) - 1.0, (rand.Float64() * 2) - 1.0},
		Id:       bid,
	}

	Boids[bid] = &b
	go b.start()
}
