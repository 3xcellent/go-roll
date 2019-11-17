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

	r := roll.ParseStrToRoll(flag.Args()[0])
	if *verbose {
		r.SetVerbose()
	}

	r.Calc()

	fmt.Println(r.String())
}
