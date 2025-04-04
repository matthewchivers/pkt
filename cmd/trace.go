package cmd

import (
	"fmt"
	"time"
)

type TraceCmd struct {
	Target string `arg:"" required:"" help:"Host to trace."`
	Watch  bool   `help:"Repeat the trace every 5 seconds." short:"w"`
}

func (t *TraceCmd) Run() error {
	if t.Watch {
		for {
			runTrace(t.Target)
			fmt.Println("---")
			time.Sleep(5 * time.Second)
		}
	} else {
		runTrace(t.Target)
	}
	return nil
}

func runTrace(target string) {
	fmt.Printf("Tracing route to %s...\n", target)
}
