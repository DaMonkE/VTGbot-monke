package rollers

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Roll deconstructs a dice description and rolls it, rolling a 1d10 if none is given
func Roll(diceStr string) (mess string) {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
	// Initialize dice variables
	diceNum := 1
	diceSides := 10
	sum := 0
	// Check which dice to roll
	if diceStr != "" {
		diceDesc := strings.Split(diceStr, "d")
		diceNum, _ = strconv.Atoi(diceDesc[0])
		diceSides, _ = strconv.Atoi(diceDesc[1])
	}
	// Initilize the roll log
	rolls := make([]int, diceNum)
	// Roll the dice!
	for i := 0; i < diceNum; i++ {
		rolls[i] = rand.Intn(diceSides) + 1
		sum += rolls[i]
	}
	// Create the result message
	mess = fmt.Sprintf(" rolled %vd%v and got %v, ( ", diceNum, diceSides, sum)
	for _, roll := range rolls {
		if roll == diceSides {
			mess += fmt.Sprintf("**%v**  ", roll)
		} else {
			mess += fmt.Sprintf("*%v  ", roll)
		}
	}
	mess += ")"
	return
}
