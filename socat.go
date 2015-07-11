package main

import (
	"fmt"
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
        sout, err := cmd.StdoutPipe()
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

	cmd.Start()

	var buf []byte
	_, err = sout.Read(buf)
	if err != nil {
                fmt.Println(err)
                os.Exit(1)
	}
        fmt.Println(string(buf))

	_, err = serr.Read(buf)
	if err != nil {
                fmt.Println(err)
                os.Exit(1)
	}
        fmt.Println(string(buf))
	cmd.Wait()
}
