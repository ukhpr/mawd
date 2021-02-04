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

// apogeeCmd represents the apogee command
var apogeeCmd = &cobra.Command{
	Use:   "apogee",
	Short: "Report current On/Off status of apogee delay",
	Long:  "Report current On/Off status of apogee delay",
	Run: func(cmd *cobra.Command, args []string) {
		showApogeeStatus()
	},
}

func init() {
	showCmd.AddCommand(apogeeCmd)
}

func showApogeeStatus() {
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

	// Report the current status of apogee delay ON/OFF
	serialport.SendCommand(port, "A\r")

	// Process response
	line := serialport.Readline(port) // Read blank line
	line = serialport.Readline(port)  // Read 1st line of response
	fmt.Println("Apogee Delay:", line)
}
