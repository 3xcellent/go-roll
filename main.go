package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/3xcellent/go-roll/roll"
)

func main() {
	verbose := flag.Bool("v", false, "verbose")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Example: roll 2d8")
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())

	formatter := roll.Simple
	if *verbose {
		formatter = roll.Verbose
	}
	r := roll.ParseStrToRoll(flag.Args()[0], formatter)

	r.Calc()

	fmt.Println(r.String())
}
