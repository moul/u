package u

import (
	"os"
	"os/signal"
)

func WaitForCtrlC() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
