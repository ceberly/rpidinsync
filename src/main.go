// Most of this code is taken from https://github.com/davecheney/gpio/tree/master/examples/blink
// Please visit Dave's repo to learn about what it's doing.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/davecheney/gpio"
	"github.com/davecheney/gpio/rpi"
)

const (
	PPQN    = 24  // Roland DIN sync spec is 24 pulses per quarter note.
	MAX_BPM = 240 // don't stress the tr-606 CPU
)

func runClock(pin gpio.Pin, bpm int) {
	if bpm > MAX_BPM {
		fmt.Printf("Please supply a reasonable tempo (< %d bpm)")
		os.Exit(0)
	}

	// in 4/4 time, 1 quarter note = 24 ppqn * (bpm / 60) pulses per second.
	// eg. 120bpm is 48 pulses per second.
	// try to make a square pulse by setting and sleeping for half a cycle each.
	hz := PPQN * (bpm / 60)
	cycle := time.Duration(1000000 / hz / 2) * time.Microsecond

	fmt.Printf("clock is running at %d bpm (%d hz with a %+v cycle)\n", bpm, hz, cycle)
	for {
			pin.Set()
			time.Sleep(cycle)
			pin.Clear()
			time.Sleep(cycle)
	}
}

func cleanup(pins []gpio.Pin) {
	for _, p := range pins {
		p.Clear()
		p.Close()
	}
}

func main() {
	clockPin, err := gpio.OpenPin(rpi.GPIO25, gpio.ModeOutput)
	if err != nil {
		fmt.Printf("Error opening pin! %s\n", err)
		os.Exit(1)
	}

	startStopPin, err := gpio.OpenPin(rpi.GPIO24, gpio.ModeOutput)
	if err != nil {
		fmt.Printf("Error opening pin! %s\n", err)
		os.Exit(1)
	}

	// Make sure pin is off on exit as good practice.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			cleanup([]gpio.Pin{clockPin, startStopPin})
			os.Exit(0)
		}
	}()

	go func() {
		runClock(clockPin, 120) // 120 bpm
	}()

	// two seconds seems to be enough for the TR-606 clock to find the right tempo
	time.Sleep(2 * time.Second)

  // run the sequencer for 10 seconds
	fmt.Println("Sequencer is running at 120bpm")
	startStopPin.Set()
	time.Sleep(10 * time.Second)
	fmt.Println("Sequencer is stopping")

	cleanup([]gpio.Pin{clockPin, startStopPin})
}
