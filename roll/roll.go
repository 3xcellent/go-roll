package roll

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

const (
	reRolls    = `^[^d]*`
	reModifier = `[\+|-]\d\d?`
	reType     = `d\d\d?`
	reHigh     = `d\d{1,2}h`
	reLow      = `d\d{1,2}l`
)

type Roll struct {
	numRolls   int
	maxScore   int
	chooseHigh bool
	chooseLow  bool
	modifier   int
}

func ParseStrToRoll(str string) Roll {
	return Roll{
		numRolls:   parseInt(str, reRolls, 1, false),
		maxScore:   parseInt(str, reType, 0, true),
		chooseHigh: parseBool(str, reHigh),
		chooseLow:  parseBool(str, reLow),
		modifier:   parseInt(str, reModifier, 0, false),
	}
}

func (r Roll) Calc() int {
	var rolls []string
	tot := 0

	for i := 0; i < r.numRolls; i++ {
		rolled := r.getRoll()
		rolls = append(rolls, fmt.Sprintf("%d", rolled))
		tot += rolled
	}

	if len(rolls) > 1 {
		fmt.Printf("rolls: %s\n", strings.Join(rolls, ","))
	}
	return tot + r.modifier
}

func (r Roll) getRoll() int {
	if r.chooseHigh {
		roll1 := rand.Intn(r.maxScore) + 1
		roll2 := rand.Intn(r.maxScore) + 1
		fmt.Printf("choosing highest of: %d, %d\n", roll1, roll2)
		if roll1 > roll2 {
			return roll1
		} else {
			return roll2
		}
	}
	if r.chooseLow {
		roll1 := rand.Intn(r.maxScore) + 1
		roll2 := rand.Intn(r.maxScore) + 1
		fmt.Printf("choosing lowest of: %d, %d\n", roll1, roll2)
		if roll1 < roll2 {
			return roll1
		} else {
			return roll2
		}
	}
	return rand.Intn(r.maxScore) + 1
}

func parseInt(argStr, reStr string, defInt int, trimBeginChar bool) int {
	re, err := regexp.Compile(reStr)
	if err != nil {
		fmt.Errorf("could not compile regex: %v", err)
	}
	argIntStr := re.FindString(argStr)

	argInt := defInt
	if argIntStr != "" {
		if trimBeginChar {
			argIntStr = argIntStr[1:]
		}
		argInt, err = strconv.Atoi(argIntStr)
		if err != nil {
			fmt.Errorf("could not parse number: %v", err)
		}
	}
	return argInt
}

func parseBool(argStr string, reStr string) bool {
	re, err := regexp.Compile(reStr)
	if err != nil {
		fmt.Errorf("could not compile regex: %v", err)
	}
	return re.MatchString(argStr)
}
