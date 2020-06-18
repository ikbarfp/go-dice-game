package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

var dice = []int{1, 2, 3, 4, 5, 6}

type Scoreboard struct {
	Player string
	Points int
}

var PairList []Scoreboard

func rollDice(arrayOfInt []int) int {
	rand.Seed(time.Now().UnixNano())

	return arrayOfInt[rand.Intn(len(arrayOfInt))]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Roll the dice Game")
	fmt.Println("---------------------")

	var playerNumber int
	var tempScore int
	var playerNames = []string{}

	fmt.Println("input amount of player : ")
	fmt.Scanf("%d", &playerNumber)

	for i := 0; i < playerNumber; i++ {
		fmt.Printf("input player-%d name : ", i+1)
		name, err := reader.ReadString('\n')
		if err != nil {
			fmt.Errorf("failed to fetch data")
			return
		}
		name = strings.Replace(name, "\n", "", -1) // if you are in unix environment (wsl not included) you can comment this code
		playerNames = append(playerNames, name)
	}

	scoreboardMap := make(map[string]int)
	for _, player := range playerNames {
		tempScore = 0
		for j := 0; j < 5; j++ {
			fmt.Printf("player %s rolled out the dice . . . ", player)
			time.Sleep(2000 * time.Millisecond)

			point := rollDice(dice)
			fmt.Printf("Got %v!\n", point)

			if point%2 == 1 {
				tempScore += 5
			} else if point%2 == 0 {
				tempScore -= 3
			}
		}
		scoreboardMap[player] = tempScore
		fmt.Printf("[%v] total score : %d\n", player, scoreboardMap[player])
	}

	for k, v := range scoreboardMap {
		PairList = append(PairList, Scoreboard{k, v})
	}

	sort.Slice(PairList, func(i, j int) bool {
		return PairList[i].Points > PairList[j].Points
	})
	fmt.Printf("\nYeay %s win the game!", PairList[0].Player)

	return
}
