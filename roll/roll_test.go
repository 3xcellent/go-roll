package roll

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestParseStrToRoll(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want Roll
	}{
		{
			name: "d20",
			args: args{"d20"},
			want: Roll{
				numRolls:       1,
				maxScore:       20,
				chooseHigh:     false,
				chooseLow:      false,
				modifier:       0,
				CalculatedRoll: 0,
				formatter:      Simple,
			},
		},
		{
			name: "2d20",
			args: args{"2d20"},
			want: Roll{
				numRolls:       2,
				maxScore:       20,
				chooseHigh:     false,
				chooseLow:      false,
				modifier:       0,
				CalculatedRoll: 0,
				formatter:      Simple,
			},
		},
		{
			name: "4d8+9",
			args: args{"4d8+9"},
			want: Roll{
				numRolls:       4,
				maxScore:       8,
				chooseHigh:     false,
				chooseLow:      false,
				modifier:       9,
				CalculatedRoll: 0,
				formatter:      Simple,
			},
		},
		{
			name: "d100",
			args: args{"d100h"},
			want: Roll{
				numRolls:       1,
				maxScore:       100,
				chooseHigh:     true,
				chooseLow:      false,
				modifier:       0,
				CalculatedRoll: 0,
				formatter:      Simple,
			},
		},
		{
			name: "d4l+8",
			args: args{"d4l+8"},
			want: Roll{
				numRolls:       1,
				maxScore:       4,
				chooseHigh:     false,
				chooseLow:      true,
				modifier:       8,
				CalculatedRoll: 0,
				formatter:      Simple,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseStrToRoll(tt.args.str)
			if !ParsedRollsAreEqual(tt.want, got) {
				t.Errorf("ParseStrToRoll()\n  want: %#v\n   got: %#v", tt.want, got)
			}
		})
	}
}

func ParsedRollsAreEqual(expected, got Roll) bool {
	return expected.numRolls == got.numRolls &&
		expected.maxScore == got.maxScore &&
		expected.chooseHigh == got.chooseHigh &&
		expected.chooseLow == got.chooseLow &&
		expected.modifier == got.modifier &&
		expected.CalculatedRoll == got.CalculatedRoll &&
		len(expected.rolls) == len(got.rolls)
}

func TestRoll_CalcRandomness(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	r := ParseStrToRoll("d20")

	totals := map[int]int{}
	for i := 1; i <= r.maxScore; i++ {
		totals[i] = 0
	}

	for i := 0; i < 1000000; i++ {
		r.Calc()
		totals[r.CalculatedRoll] = totals[r.CalculatedRoll] + 1
	}

	for i := 1; i <= r.maxScore; i++ {
		fmt.Printf("%d - %d\n", i, totals[i])
	}
}
