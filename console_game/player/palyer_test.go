package player

import (
	"testing"
)

func TestDigHole(t *testing.T) {
	player := New()
	intense := true
	player.DigHole(intense)
	got := player
	want := Player{
		burrowLength: 15,
		health:       70,
		respect:      20,
		weight:       30,
	}
	if *got != want {
		t.Errorf("got %+v, wanted %+v", *got, want)
	}

	player.DigHole(!intense)
	want.burrowLength += 2
	want.health -= 10
	if *got != want {
		t.Errorf("got %+v, wanted %+v", *got, want)
	}
}

func TestEatGrass(t *testing.T) {
	player := New()
	green := true
	got := player
	got.respect = 31
	player.EatGrass(green)
	want := Player{
		burrowLength: 10,
		health:       130,
		respect:      31,
		weight:       60,
	}
	if *got != want {
		t.Errorf("got %+v, wanted %+v", *got, want)
	}

	got.respect, want.respect = 29, 29
	want.health -= 30
	player.EatGrass(green)
	if *got != want {
		t.Errorf("got %+v, wanted %+v", *got, want)
	}

	got.respect, want.respect = 31, 31
	want.health += 30
	want.weight += 15
	player.EatGrass(!green)
	if *got != want {
		t.Errorf("got %+v, wanted %+v", *got, want)
	}
}
