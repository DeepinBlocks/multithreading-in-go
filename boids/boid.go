package boids

import (
	"math"
	"math/rand"
	"time"
)

const (
	ScreenWidth, ScreenHeight = 640, 360
	BoidCount                 = 500
	ViewRadius                = 13
	AdjRate                   = 0.015
)

var (
	Boids   [BoidCount]*Boid
	BoidMap [ScreenWidth + 1][ScreenHeight + 1]int
)

type Boid struct {
	Position Vector2D
	Velocity Vector2D
	Id       int
}

func (b *Boid) calcAcceleration() Vector2D {
	upper, lower := b.Position.AddV(ViewRadius), b.Position.AddV(-ViewRadius)
	avgVelocity := Vector2D{0, 0}
	count := 0.0
	for i := math.Max(lower.X, 0); i <= math.Min(upper.X, ScreenWidth); i++ {
		for j := math.Max(lower.Y, 0); j <= math.Min(upper.Y, ScreenHeight); j++ {
			if otherBoidId := BoidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.Id {
				if dist := Boids[otherBoidId].Position.Distance(b.Position); dist < ViewRadius {
					count++
					avgVelocity = avgVelocity.Add(Boids[otherBoidId].Velocity)
				}
			}
		}
	}
	accel := Vector2D{0, 0}
	if count > 0 {
		avgVelocity = avgVelocity.DivisionV(count)
		accel = avgVelocity.Subtract(b.Velocity).MultiplyV(AdjRate)
	}
	return accel
}

func (b *Boid) moveOne() {
	b.Velocity = b.Velocity.Add(b.calcAcceleration()).limit(-1, 1)
	BoidMap[int(b.Position.X)][int(b.Position.Y)] = -1
	b.Position = b.Position.Add(b.Velocity)
	BoidMap[int(b.Position.X)][int(b.Position.Y)] = b.Id

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
	BoidMap[int(b.Position.X)][int(b.Position.Y)] = b.Id

	go b.start()
}
