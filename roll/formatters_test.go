package roll

import "testing"

func TestSimple(t *testing.T) {
	type args struct {
		r Roll
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "7", args: args{Roll{CalculatedRoll: 7}}, want: "7"},
		{name: "19", args: args{Roll{CalculatedRoll: 19}}, want: "19"},
		{name: "100", args: args{Roll{CalculatedRoll: 100}}, want: "100"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Simple(tt.args.r); got != tt.want {
				t.Errorf("Simple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerbose(t *testing.T) {
	type args struct {
		r Roll
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "single roll does not show individual rolls",
			args: args{
				Roll{
					numRolls:       1,
					maxScore:       20,
					chooseHigh:     false,
					chooseLow:      false,
					modifier:       0,
					CalculatedRoll: 9,
					rolls:          []int{4},
				},
			},
			want: "Roll: 9 (min/max 1/20)",
		},
		{
			name: "multiple rolls show individual rolls",
			args: args{
				Roll{
					numRolls:       4,
					maxScore:       8,
					chooseHigh:     false,
					chooseLow:      false,
					modifier:       0,
					CalculatedRoll: 13,
					rolls:          []int{4, 5, 3, 1},
				},
			},
			want: "Rolls: 4, 5, 3, 1\nTotal: 13 (min/max 4/32)",
		},
		{
			name: "Modifiers show correctly for positive numbers",
			args: args{
				Roll{
					numRolls:       1,
					maxScore:       8,
					chooseHigh:     false,
					chooseLow:      false,
					modifier:       9,
					CalculatedRoll: 22,
					rolls:          []int{4},
				},
			},
			want: "Modifier: +9\nTotal: 22 (min/max 10/17)",
		},
		{
			name: "Modifiers show correctly for negative numbers",
			args: args{
				Roll{
					numRolls:       1,
					maxScore:       8,
					chooseHigh:     false,
					chooseLow:      false,
					modifier:       -2,
					CalculatedRoll: 2,
					rolls:          []int{4},
				},
			},
			want: "Modifier: -2\nTotal: 2 (min/max -1/6)",
		},
		{
			name: "d20h - keeping highest number performs the rolls twice",
			args: args{
				Roll{
					numRolls:       1,
					maxScore:       20,
					chooseHigh:     true,
					chooseLow:      false,
					modifier:       0,
					CalculatedRoll: 15,
					rolls:          []int{3, 15},
				},
			},
			want: "Rolls: (3, 15)\nTotal: 15 (min/max 1/20)",
		},
		{
			name: "keeping highest number performs the rolls twice",
			args: args{
				Roll{
					numRolls:       2,
					maxScore:       8,
					chooseHigh:     true,
					chooseLow:      false,
					modifier:       0,
					CalculatedRoll: 12,
					rolls:          []int{4, 2, 6, 8},
				},
			},
			want: "Rolls: (4, 2), (6, 8)\nTotal: 12 (min/max 2/16)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Verbose(tt.args.r); got != tt.want {
				t.Errorf("Simple()\nwanted:\n%v\ngot:\n%v", tt.want, got)
			}
		})
	}
}
