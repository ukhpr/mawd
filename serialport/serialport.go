package serialport

import (
	"fmt"
	"log"

	"go.bug.st/serial"
)

// IsValid checks if the selected Serial Port is detected as a valid port
func IsValid(port string) error {
	ports, err := serial.GetPortsList()
	if err != nil {
		return err
	}
	if len(ports) == 0 {
		return fmt.Errorf("No serial ports found")
	}
	for _, p := range ports {
		if p == port {
			return nil
		}
	}
	return fmt.Errorf("Serial port not detected")
}

// SendCommand sends an instruction string to the altimeter
func SendCommand(p serial.Port, cmd string) {
	p.Write([]byte(cmd))
}

// Readline returns a complete line read from the serial port terminated with \n
func Readline(p serial.Port) string {
	buff := make([]byte, 1)
	line := ""
	for {
		// Reads up to 100 bytes
		n, err := p.Read(buff)
		if err != nil {
			log.Fatal(err)
			break
		}
		if n == 0 || string(buff[:n]) == "\n" {
			break
		}
		line += string(buff[:n])
	}
	return line
}
