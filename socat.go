package main

import (
	"fmt"
	"bufio"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("socat", "-d", "-d", "pty,raw,echo=0", "pty,raw,echo=0")
	serr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd.Start()

	scanner := bufio.NewScanner(serr)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		fmt.Println()
	}	

	cmd.Wait()
}
