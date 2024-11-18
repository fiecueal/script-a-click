package main

// robotgo docs: https://pkg.go.dev/github.com/go-vgo/robotgo
import (
	"flag"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/robotn/gohook"
)

func main() {
	interval := flag.Int("interval", 1000, "The time between the previous click and the next click")
	button := flag.String("button", robotgo.Mleft, "The mouse button to simulate clicking")
	repetitions := flag.Int("repetitions", -1, "The number of times the button is clicked")
	flag.Parse()

	// nonblocking click loop
	go func() {
		for i := 0; *repetitions < 0 || i < *repetitions; i++ {
			time.Sleep(time.Duration(*interval) * time.Millisecond)
			robotgo.Click(*button)
		}
		hook.End()
		os.Exit(0)
	}()

	// start hook listener
	s := hook.Start()

	// kill program on any key pressed
	hook.Register(hook.KeyDown, []string{}, func(e hook.Event) {
		hook.End()
		os.Exit(0)
	})

	// process hook event
	<-hook.Process(s)
}
