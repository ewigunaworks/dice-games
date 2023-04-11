package main

import (
	"bufio"
	"fmt"
	"os"
	"qoin/src/service"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input Player")
	fmt.Println("---------------------")
	player, _ := reader.ReadString('\n')
	fmt.Println("---------------------")

	fmt.Println("Input Dice")
	fmt.Println("---------------------")
	dice, _ := reader.ReadString('\n')
	fmt.Println("---------------------")

	playerTrim := strings.TrimSuffix(player, "\n")
	diceTrim := strings.TrimSuffix(dice, "\n")

	playerInt, _ := strconv.Atoi(playerTrim)
	diceInt, _ := strconv.Atoi(diceTrim)

	_, err := service.NewGame(playerInt, diceInt)
	if err != nil {
		fmt.Println(err)
	}
}
