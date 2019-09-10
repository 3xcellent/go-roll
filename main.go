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

	r := roll.ParseStrToRoll(os.Args[1], rand.New(rand.NewSource(time.Now().Unix())))
	fmt.Println(r.Calc())
}
