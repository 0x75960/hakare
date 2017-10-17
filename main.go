package main

import (
	"flag"
	"fmt"
	"os/exec"
	"time"
)

func init() {
}

func main() {
	flag.Parse()
	cl := flag.Args()

	cmd := exec.Command(cl[0], cl[1:]...)

	start := time.Now()
	cmd.Run()
	end := time.Now()

	fmt.Println("It takes", (end.Sub(start)).Seconds(), "sec.")
}
