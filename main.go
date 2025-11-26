package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type Player struct {
    Number int
    Value int
}


type Game struct {
	Players []Player
	UserChoice int
}


/*
Комбинации победы

1 - камень
2 - ножницы
3 - бумага
*/
var winComb = map[int]int{
    3: 1,
    1: 2,
	2: 3,
}


func getRandomValForBot() int {
	return rand.Intn(3) + 1
}


func getPlayValues(game Game) []int {
	allValues := make([]int, 0, len(game.Players)+1)
	for _, p := range game.Players {
        allValues = append(allValues, p.Value)
	}
	allValues = append(allValues, game.UserChoice)
	return allValues
}

func gameCore(game Game) (*int, error) {
	playValues := getPlayValues(game)
    if len(playValues) < 2 {
        return nil, errors.New("недостаточное количество игроков")
    }

	win_choice := playValues[0]
	one_comb := true
	for i := 1; i < len(playValues); i++ {
		v := playValues[i]
		if v == win_choice {
			continue
		}
		if winComb[v] == win_choice {
			if !one_comb {
				return nil, nil
			}
			win_choice = v
		}
		one_comb = false
	}
	if one_comb {
		return nil, nil
	}
	return &win_choice, nil
}

func getUserResult(game Game, winValue int) string {
	if game.UserChoice == winValue {
		return "Вы победили"
	}
	return "Вы проиграли"
}

func main() {
	var numBots int
	var userVal int


	for {
		fmt.Print("Введите количество ботов (1 - 6):\n")
		_, err := fmt.Scanln(&numBots)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}
		if numBots < 1 || numBots > 6 {
			fmt.Println("Количество ботов должно быть от 1 до 6")
		} else { break }
    }

	for {
		fmt.Print("Введите значение (1 - камень, 2 - ножницы, 3 - бумага):\n")
		fmt.Scanf("%d", &userVal)
		if userVal < 1 || userVal > 3 {
			fmt.Println("Значение должно быть от 1 до 3")
			continue
		} else {
			bots := make([]Player, numBots)
			for i := 0; i < numBots; i++ {
				bots[i] = Player{
					Number: i+1,
					Value: getRandomValForBot(),
				}
			}
		
			game := Game{
				Players: bots,
				UserChoice: userVal,
			}
		
			fmt.Println(game)
		
			winValue, err := gameCore(game)
			if err != nil {
				fmt.Println("Ошибка:", err)
				return
			} else if winValue == nil {
				fmt.Println("Ничья")
				continue
			} else {
				fmt.Println("Выбор победителя:", *winValue)
				userResult := getUserResult(game, *winValue)
				fmt.Println(userResult)
			}
		}
	}
}