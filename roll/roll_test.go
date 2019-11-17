package roll

import (
	"fmt"
	"math"
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

func TestRoll_Calc(t *testing.T) {
	t.Run("sets correct number of rolls", func(t *testing.T) {
		type fields struct {
			numRolls       int
			maxScore       int
			chooseHigh     bool
			chooseLow      bool
			modifier       int
			CalculatedRoll int
			rolls          []int
			formatter      OutputFormatter
		}
		tests := []struct {
			name          string
			parseString   string
			expRollLength int
		}{
			{
				name:          "for one roll",
				parseString:   "d20",
				expRollLength: 1,
			},
			{
				name:          "for two rolls",
				parseString:   "2d20",
				expRollLength: 2,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				r := ParseStrToRoll(tt.parseString)
				r.Calc()
				rollCount := len(r.rolls)
				if rollCount != tt.expRollLength {
					t.Errorf("Calc() expected # rows: %d | got : %d", tt.expRollLength, rollCount)
				}
			})
		}
	})
}

func TestRoll_CalcRandomness(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	r := ParseStrToRoll("d20")
	testSize := int(math.Pow(2, 24)) // ~ 16 mil
	fmt.Printf("test size: %d", testSize)
	totals := map[int]int{}
	for i := 1; i <= r.maxScore; i++ {
		totals[i] = 0
	}

	for i := 0; i < testSize; i++ {
		r.Calc()
		totals[r.CalculatedRoll] = totals[r.CalculatedRoll] + 1
	}

	subGroupSize := testSize / r.maxScore
	subGroupDeviations := make([]int, r.maxScore)

	for i := 1; i <= r.maxScore; i++ {
		subGroupDeviations[i-1] = absInt(subGroupSize - totals[i])
	}
	avg, min, max := avgMinMaxDeviations(subGroupDeviations)
	fmt.Printf("test size: %d | avg +-: %d (%f%%) | min: %d (%f%%)  | max: %d (%f%%)\n",
		testSize,
		avg,
		float64(avg)/float64(subGroupSize),
		min,
		float64(min)/float64(subGroupSize),
		max,
		float64(max)/float64(subGroupSize),
	)
}

func avgMinMaxDeviations(slc []int) (int, int, int) {
	tot, min, max := 0, slc[0], slc[0]
	for _, d := range slc {
		if d > max {
			max = d
		}
		if d < min {
			min = d
		}
		tot += d
	}

	return tot / len(slc), min, max
}

func absInt(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
