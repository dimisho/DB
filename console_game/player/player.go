package player

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	burrowLength int
	health       int
	respect      int
	weight       int
}

const (
	rewardForWeakCreature   = 10
	rewardForMediumCreature = 20
	rewardForStrongCreature = 40
)

func New() *Player {
	return &Player{
		burrowLength: 10,
		health:       100,
		respect:      20,
		weight:       30,
	}
}

func (player *Player) OutputCharacteristics() string {
	return fmt.Sprintf("\nТекущие характеристики:\nДлина норы: %v\nЗдоровье: %v\nУважение: %v\nВес: %v",
		player.burrowLength,
		player.health,
		player.respect,
		player.weight,
	)
}

func (player *Player) Sleep() {
	player.burrowLength -= 2
	player.health += 20
	player.respect -= 2
	player.weight -= 5
}

func (player *Player) DigHole(intense bool) {
	if intense {
		player.burrowLength += 5
		player.health -= 30
	} else {
		player.burrowLength += 2
		player.health -= 10
	}
}

func (player *Player) EatGrass(green bool) {
	if green {
		if player.respect < 30 {
			player.health -= 30
		} else {
			player.health += 30
			player.weight += 30
		}
	} else {
		player.health += 30
		player.weight += 15
	}
}

func (player *Player) Fight(level int) {
	probability := player.weight / (level + player.weight)
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	chance := generator.Float64()

	switch {
	case float64(probability) > 0.5:
		if chance >= float64(probability) {
			player.respect += rewardForWeakCreature
		} else {
			player.health -= rewardForWeakCreature
		}
	case float64(probability) == 0.5:
		if chance >= float64(probability) {
			player.respect += rewardForMediumCreature
		} else {
			player.health -= rewardForMediumCreature
		}
	default:
		if chance >= float64(probability) {
			player.respect += rewardForStrongCreature
		} else {
			player.health -= rewardForStrongCreature
		}
	}
}

func (player *Player) IsLose() bool {
	return player.burrowLength <= 0 || player.health <= 0 || player.respect <= 0 || player.weight <= 0
}

func (player *Player) IsWin() bool {
	return player.respect > 100
}
