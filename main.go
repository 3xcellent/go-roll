package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/3xcellent/go-roll/roll"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Example: roll 2d8")
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())
	r := roll.ParseStrToRoll(os.Args[1])
	r.Calc()
	fmt.Println(r.String())
}
