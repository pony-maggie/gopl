// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 244.

// Countdown implements the countdown for a rocket launch.
package main

import (
	"fmt"
	"os"
	"time"
)

//!+

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("press return to abort")
	select {
	case <-time.After(10 * time.Second):
	case <-abort:
		fmt.Println("launch aborted!")
		return
	}

	launch()
}

//!-

func launch() {
	fmt.Println("Lift off!")
}
