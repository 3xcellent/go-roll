package roll

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"

	"github.com/sirupsen/logrus"
)

const (
	reRolls    = `^[^d]*`
	reModifier = `[\+|-]\d{1,2}`
	reType     = `d\d{1,3}`
	reHigh     = `d\d{1,3}h`
	reLow      = `d\d{1,3}l`
)

type OutputFormatter func(Roll) string

type Roll struct {
	numRolls       int
	maxScore       int
	chooseHigh     bool
	chooseLow      bool
	modifier       int
	CalculatedRoll int
	rolls          []int
	formatter      OutputFormatter
}

func ParseStrToRoll(str string) Roll {
	return Roll{
		numRolls:   parseInt(str, reRolls, 1, false),
		maxScore:   parseInt(str, reType, 0, true),
		chooseHigh: parseBool(str, reHigh),
		chooseLow:  parseBool(str, reLow),
		modifier:   parseInt(str, reModifier, 0, false),
		formatter:  Simple,
	}
}

func (r *Roll) SetVerbose() {
	r.formatter = Verbose
}

func (r *Roll) Calc() {
	tot := 0

	for i := 0; i < r.numRolls; i++ {
		if r.chooseHigh || r.chooseLow {
			roll1 := r.getRoll()
			roll2 := r.getRoll()
			r.rolls = append(r.rolls, roll1, roll2)
			if r.chooseHigh {
				if roll1 > roll2 {
					tot += roll1
				} else {
					tot += roll2
				}
			} else {
				if roll1 < roll2 {
					tot += roll1
				} else {
					tot += roll2
				}
			}
		} else {
			rolled := r.getRoll()
			r.rolls = append(r.rolls, rolled)
			tot += rolled
		}
		logrus.Debugf("roll: %d, total: %d | rolls: %+v", i+1, tot, r.rolls)
	}

	r.CalculatedRoll = tot + r.modifier
}

func (r *Roll) String() string {
	return r.formatter(*r)
}

func (r *Roll) getRoll() int {
	return rand.Intn(r.maxScore) + 1
}

func parseInt(argStr, reStr string, defInt int, trimBeginChar bool) int {
	re, err := regexp.Compile(reStr)
	if err != nil {
		fmt.Printf("could not compile regex: %v", err)
	}
	argIntStr := re.FindString(argStr)

	argInt := defInt
	if argIntStr != "" {
		if trimBeginChar {
			argIntStr = argIntStr[1:]
		}
		argInt, err = strconv.Atoi(argIntStr)
		if err != nil {
			fmt.Printf("could not parse number: %v", err)
		}
	}
	return argInt
}

func parseBool(argStr string, reStr string) bool {
	re, err := regexp.Compile(reStr)
	if err != nil {
		fmt.Printf("could not compile regex: %v", err)
	}
	return re.MatchString(argStr)
}
