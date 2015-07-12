package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
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
		re := regexp.MustCompile("([/]dev[/]ttys[0-9]{3})")
		line := []byte(scanner.Text())
		if re.Match(line) {
			fmt.Println(string(re.Find(line)))
		}
	}	

	cmd.Wait()
}
