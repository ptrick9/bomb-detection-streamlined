package main

import (
	"testing"
	"math"
)


func TestDist(t *testing.T) {
	n := bn{NodeImpl:&NodeImpl{x: 3, y: 3, id: 0}, x_speed: rangeInt(-3, 3), y_speed: rangeInt(-3, 3)}
	b := bomb{1,1}

	dist := n.distance(b)
	if dist != float32(math.Sqrt(8)) {
		t.Errorf("Distance was incorrect, got: %f, want: %f", dist, float32(math.Sqrt(8)))
	}

}

func TestBnMove(t *testing.T) {
	n := bn{NodeImpl:&NodeImpl{x: 3, y: 3, id: 0}, x_speed: -4, y_speed: -4}
	n.move()

	if n.x != 1 && n.y != 1 {
		t.Errorf("BN Movement Incorrect. Expected %d %d, got %d %d", 1, 1, n.x, n.y)
	}
}


func TestWnMove(t *testing.T) {
	n := wn{NodeImpl:&NodeImpl{x: 3, y: 3, id: 0}, speed: -4, dir: 0}
	n.move()

	if n.x != 1 && n.y != 3 {
		t.Errorf("WN Movement Incorrect. Expected %d %d, got %d %d", 1, 3, n.x, n.y)
	}
}