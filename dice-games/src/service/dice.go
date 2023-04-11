package service

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type DicePlayer struct {
	player int
	dice   []int
	point  int
	status string
}

func NewGame(player int, dice int) ([]DicePlayer, error) {

	if player <= 1 {
		return nil, errors.New("Need another player to play the game")
	}

	if dice <= 1 {
		return nil, errors.New("Need dice roll to play the game")
	}

	arrPlayer := NumberOfPlayer(player)
	dicePlayer := []DicePlayer{}
	for _, v := range arrPlayer {
		dicePerPlayer := NumberOfDicePerPlayer(dice)
		dicePlayer = append(dicePlayer, DicePlayer{
			player: v,
			dice:   dicePerPlayer,
			point:  0,
			status: "playing",
		})
	}

	fmt.Printf("Pemain: %d, Dadu: %d\n", player, dice)
	fmt.Println("---------------------")

	StartGame(dicePlayer)

	sort.Slice(dicePlayer, func(i, j int) bool {
		return dicePlayer[i].point > dicePlayer[j].point
	})

	getWinner(dicePlayer)

	return dicePlayer, nil
}

func NumberOfPlayer(player int) []int {
	var arrPlayer []int
	for i := 1; i <= player; i++ {
		arrPlayer = append(arrPlayer, i)
	}

	return arrPlayer
}

func NumberOfDicePerPlayer(dice int) []int {
	var dicePerPlayer []int
	min := 1
	max := 7
	for i := 0; i < dice; i++ {
		rand.Seed(time.Now().UnixNano())
		dicePerPlayer = append(dicePerPlayer, rand.Intn(max-min)+min)
	}

	return dicePerPlayer
}

func StartGame(playing []DicePlayer) {
	i := 0
	turnCount := 1
	for {
		countDicePlayer := 0
		fmt.Printf("Giliran %d lempar dadu:\n", turnCount)
		// // loop through playing to get before evaluate
		getPlayerDetail(playing)

		// var tempArr []int
		tempCount := 0
		for idx, _ := range playing {
			// evaluate
			for i := len(playing[idx].dice) - (1 + tempCount); i >= 0; i-- {
				// check if the dice have 6, add idx to temp array
				if playing[idx].dice[i] == 6 || playing[idx].dice[i] == 1 {
					if playing[idx].dice[i] == 6 {
						// add point + 1
						playing[idx].point += 1
					} else if playing[idx].dice[i] == 1 {
						tempCount++
						nextIdx := idx + 1
						// check if already in last idx
						if idx == len(playing)-1 {
							nextIdx = 0
						}
						// append new item to next arr
						if len(playing[nextIdx].dice) == 0 {
							if nextIdx == len(playing)-1 {
								nextIdx = 0
							} else {
								nextIdx += 1
							}
							playing[nextIdx].dice = append(playing[nextIdx].dice, playing[idx].dice[i])
						} else {
							playing[nextIdx].dice = append(playing[nextIdx].dice, playing[idx].dice[i])
						}
					}
					playing[idx].dice = append(playing[idx].dice[:i], playing[idx].dice[i+1:]...)
				}

				// check if the out of dice, set status to finish
				if len(playing[idx].dice) < 1 {
					playing[idx].status = "finished"
				}
			}
		}

		fmt.Printf("Setelah evaluasi:\n")
		// loop through playing to get before evaluate
		getPlayerDetail(playing)
		fmt.Println("---------------------")

		// count remaining player
		for k, _ := range playing {
			if len(playing[k].dice) >= 1 {
				countDicePlayer++
			}
		}

		// // generate new dice
		for idx, _ := range playing {
			playing[idx].dice = NumberOfDicePerPlayer(len(playing[idx].dice))
		}
		i++
		turnCount++

		if i == len(playing) {
			i = 0
		}

		if countDicePlayer <= 1 {
			break
		}
	}
}

func getPlayerDetail(playing []DicePlayer) {
	for j, _ := range playing {
		if len(playing[j].dice) < 1 {
			fmt.Printf("Pemain #%d (%d): _ (Berhenti bermain karena tidak memiliki dadu)\n", j+1, playing[j].point)
		} else {
			fmt.Printf("Pemain #%d (%d): %v\n", j+1, playing[j].point, playing[j].dice)
		}
	}
}

func getWinner(playing []DicePlayer) {
	winner := 0
	lose := 0
	highscore := 0
	for i, _ := range playing {
		if playing[i].point > highscore && playing[i].status == "finished" && len(playing[i].dice) < 1 {
			winner = playing[i].player
			highscore = playing[i].point
		}

		if len(playing[i].dice) >= 1 && playing[i].status == "playing" {
			lose = playing[i].player
		}
	}
	fmt.Printf("Game berakhir karena hanya pemain #%d yang memiliki dadu.\n", lose)
	fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya.", winner)
}
