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

// telemetryCmd represents the telemetry command
var telemetryCmd = &cobra.Command{
	Use:   "telemetry",
	Short: "Report current On/Off status of live telemetry output",
	Long:  "Report current On/Off status of live telemetry output",
	Run: func(cmd *cobra.Command, args []string) {
		showTelemetryStatus()
	},
}

func init() {
	showCmd.AddCommand(telemetryCmd)
}
func showTelemetryStatus() {
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

	// Report the live telemetry setting on the device
	serialport.SendCommand(port, "T\r")

	// Process response
	_ = serialport.Readline(port)     // Read and discard blank line
	line := serialport.Readline(port) // Read 1st line of response
	fmt.Println("Live Telemetry Status:", line)
}
