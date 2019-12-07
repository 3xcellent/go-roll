package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/3xcellent/go-roll/roll"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Println("\tgo-roll 2d8 [-v]")
		fmt.Println()
		fmt.Println("\tgo-roll d20h (roll a d20 twice but only keep high; for example: with advantage)")
		fmt.Println("\tgo-roll d20l (roll a d20 twice but only keep low; for example: with disadvantage)")
		fmt.Println("\tgo-roll 2d6+7 (roll two d6 and add +7 modifier)")
		fmt.Println()
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	}
}

func main() {
	verbose := flag.Bool("v", false, "output verbose roll information")
	debug := flag.Bool("d", false, "outout debug information")
	flag.Parse()

	if *debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.ErrorLevel)
	}

	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(0)
	}

	rand.Seed(time.Now().UnixNano())

	r := roll.ParseStrToRoll(flag.Args()[0])
	if *verbose {
		r.SetVerbose()
	}

	r.Calc()

	fmt.Println(r.String())
}
