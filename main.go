package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gen2brain/beeep"
)

var usage = `
Usage: timer [duration]

Examples:
  timer 10s # 10 seconds
  timer 1m  # 1 minute
  timer 1h  # 1 hour
`

type secondsTimer struct {
	timer *time.Timer
	end   time.Time
}

func newSecondsTimer(d time.Duration) *secondsTimer {
	return &secondsTimer{
		timer: time.NewTimer(d),
		end:   time.Now().Add(d),
	}
}

func main() {
	flag.Usage = func() {
		println(usage)
	}
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		return
	}

	d, err := time.ParseDuration(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "timer: invalid duration %q, msg: %s\n", flag.Arg(0), err)
		return
	}

	stimer := newSecondsTimer(d)
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-stimer.timer.C:
			beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
			fmt.Println("\nTime's up!")
			return
		case <-ticker.C:
			fmt.Printf("\rremaining: %v", time.Until(stimer.end).Round(time.Second))
		}
	}
}
