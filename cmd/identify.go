/*
Copyright © 2020 Andrew Moore

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/ukhpr/mawd/serialport"
	"go.bug.st/serial"
)

// identifyCmd represents the identify command
var identifyCmd = &cobra.Command{
	Use:   "identify",
	Short: "Identify altimeter model",
	Run: func(cmd *cobra.Command, args []string) {
		identify()
	},
}

func init() {
	rootCmd.AddCommand(identifyCmd)
}

func identify() {
	// Open serial port
	mode := &serial.Mode{
		BaudRate: 38400,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	port, err := serial.Open(port, mode)
	if err != nil {
		log.Fatal(err)
	}

	// Request the altimeter model information
	serialport.SendCommand(port, "I")

	// Process response
	_ = serialport.Readline(port)     // Read and discard initial blank line
	line := serialport.Readline(port) // Read 1st line of response
	fmt.Println(line)
	line = serialport.Readline(port) // Read 2nd line of response
	fmt.Println(line)
}
