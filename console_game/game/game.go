package game

import (
	"console_game/store"
	"fmt"
	"io"
	"os"
)

type Game struct {
	player *store.Player
	reader io.Reader
	writer io.Writer
}

const (
	DigHole  = 1
	EatGrass = 2
	Fight    = 3
	Sleep    = 4
)

const (
	weightOfWeakCreature   = 30
	weightOfMediumCreature = 50
	weightOfStrongCreature = 70
)

func NewGame(reader io.Reader, writer io.Writer) *Game {
	return &Game{store.NewPlayer(), reader, writer}
}

func StartGame() {
	var action int
	var subAction int
	game := NewGame(os.Stdin, os.Stdout)

	fmt.Fprint(game.writer, "Console game by Shust0vD\n")
	fmt.Fprint(game.writer, "Нажмите любую клавишу чтобы начать...\n")
	fmt.Scanln()

	for {
		fmt.Fprintln(game.writer, game.player.OutputCharacteristics())
		fmt.Fprint(game.writer, "\nВыберите дейтсвие:\n1. Копать нору\n2. Поесть травки\n3. Подраться\n4. Поспать\n")
		for {
			if fmt.Fscanln(game.reader, &action); action < 1 || action > 4 {
				fmt.Fprint(game.writer, "Нет такого варианта действий, попробуйте еще раз\n")
			} else {
				break
			}
		}
		if action == DigHole {
			fmt.Fprint(game.writer, "Выберите, как вы будете копать:\n1. Интенсивно\n2. Лениво\n")
			for {
				if fmt.Fscanln(game.reader, &subAction); subAction < 1 || subAction > 2 {
					fmt.Fprint(game.writer, "Нет такого варианта действий, попробуйте еще раз\n")
				} else {
					break
				}
			}
			game.performAction(action, subAction == 1)
		} else if action == EatGrass {
			fmt.Fprint(game.writer, "Выберите, какую траву поесть:\n1. Зеленую\n2. Чухлую\n")
			for {
				if fmt.Fscanln(game.reader, &subAction); subAction < 1 || subAction > 2 {
					fmt.Fprint(game.writer, "Нет такого варианта действий, попробуйте еще раз\n")
				} else {
					break
				}
			}
			game.performAction(action, subAction == 1)
		} else {
			game.performAction(action)
		}

		if game.player.IsLose() {
			fmt.Fprint(game.writer, "Характеристики упали до 0, вы проиграли...\n")
			break
		} else if game.player.IsWin() {
			fmt.Fprint(game.writer, "Уважение больше 100, это победа!!!\n")
			break
		}

		game.player.Sleep()
		fmt.Fprintln(game.writer, "\nПрошла ночь, характеристики именились.")

		if game.player.IsLose() {
			fmt.Fprint(game.writer, "Характеристики упали до 0, вы проиграли...")
			break
		}
	}
	fmt.Fprint(game.writer, "Игра окончена")
	os.Exit(0)
}

func (game Game) performAction(action int, subAction ...bool) {
	switch action {
	case DigHole:
		game.player.DigHole(subAction[0])
	case EatGrass:
		game.player.EatGrass(subAction[0])
	case Sleep:
		game.player.Sleep()
	case Fight:
		fmt.Fprint(game.writer, "Выберите силу противника:\n1: Слабый\n2: Средний\n3: Сильный\n")
		var choiceOpponent int
		for {
			if fmt.Fscan(game.reader, &choiceOpponent); choiceOpponent < 0 || choiceOpponent > 3 {
				fmt.Fprint(game.writer, "Нет такого варианта действий, попробуйте еще раз\n")
			} else {
				break
			}
		}

		switch choiceOpponent {
		case 1:
			game.player.Fight(weightOfWeakCreature)
		case 2:
			game.player.Fight(weightOfMediumCreature)
		case 3:
			game.player.Fight(weightOfStrongCreature)
		}

	default:
		fmt.Fprint(game.writer, "Нет такого действия")
	}
}
