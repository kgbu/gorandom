package gorandom

import (
	"fmt"
	"log"
	"github.com/tarm/goserial"
)

func Act() int {
	return 3
}


func main() {
        serialConfig := &serial.Config{Name: "/dev/ttyusbserial", Baud: 9600}
        serialPort, err := serial.OpenPort(serialConfig)
        if err != nil {
                fmt.Println(serialConfig, serialPort, err)
                log.Fatal("cannot open serial port")
        }
}
