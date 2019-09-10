package roll

import (
	"fmt"
	"math/rand"
	"reflect"
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
				numRolls:   1,
				maxScore:   20,
				chooseHigh: false,
				chooseLow:  false,
				modifier:   0,
			},
		},
		{
			name: "2d20",
			args: args{"2d20"},
			want: Roll{
				numRolls:   2,
				maxScore:   20,
				chooseHigh: false,
				chooseLow:  false,
				modifier:   0,
			},
		},
		{
			name: "4d8+9",
			args: args{"4d8+9"},
			want: Roll{
				numRolls:   4,
				maxScore:   8,
				chooseHigh: false,
				chooseLow:  false,
				modifier:   9,
			},
		},
		{
			name: "2d4-1",
			args: args{"2d4h-1"},
			want: Roll{
				numRolls:   2,
				maxScore:   4,
				chooseHigh: true,
				chooseLow:  false,
				modifier:   -1,
			},
		},
		{
			name: "d4l+8",
			args: args{"d4l+8"},
			want: Roll{
				numRolls:   1,
				maxScore:   4,
				chooseHigh: false,
				chooseLow:  true,
				modifier:   8,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseStrToRoll(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseStrToRoll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoll_Calc(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	r := ParseStrToRoll("d20")

	totals := map[int]int{}
	for i := 1; i <= r.maxScore; i++ {
		totals[i] = 0
	}

	for i := 0; i < 10000; i++ {
		val := r.Calc()
		totals[val] = totals[val] + 1
	}

	for i := 1; i <= 20; i++ {
		fmt.Printf("%d - %d\n", i, totals[i])
	}
}
