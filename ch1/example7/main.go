// Hangling operating system signals
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Creat the channel where the received signal would be sent.
	// The Notify will not flock when the signal is sent and the
	// chennel is not ready.
	sChan := make(chan os.Signal, 1)

	// Notify all catch the given signals and send the os.Signal value through the sChan.
	// If no signal specified in argument, all signals are matched.
	signal.Notify(sChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// Create channel to wait till the singal is handled.
	exitChan := make(chan int)

	go func() {
		signal := <-sChan
		switch signal {
		case syscall.SIGHUP:
			fmt.Println("The calling terminal has been close")
			exitChan <- 0
		case syscall.SIGINT:
			fmt.Println("The process has been interrupted by CTRL-C")
			exitChan <- 1

		case syscall.SIGTERM:
			fmt.Println("kill SIGTERM was executed for process")
			exitChan <- 1

		case syscall.SIGQUIT:
			fmt.Println("kill SIGQUIT was executed for process")
			exitChan <- 1
		}
	}()

	code := <-exitChan
	os.Exit(code)
}
