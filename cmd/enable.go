/*
Copyright © 2021 Andrew Moore

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
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/ukhpr/mawd/serialport"
	"go.bug.st/serial"
)

// enableCmd represents the enable command
var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable function on altimeter",
	Long:  `Enable either apogee delay, continuity sequence or telemetry output`,
	Run: func(cmd *cobra.Command, args []string) {
		continuity, _ := cmd.Flags().GetBool("continuity")
		apogee, _ := cmd.Flags().GetBool("apogee")
		telemetry, _ := cmd.Flags().GetBool("telemetry")
		enable(continuity, apogee, telemetry)
	},
}

func init() {
	rootCmd.AddCommand(enableCmd)

	enableCmd.Flags().BoolP("continuity", "c", false, "enable continuity beep")
	enableCmd.Flags().BoolP("apogee", "a", false, "enable 1 second apogee delay")
	enableCmd.Flags().BoolP("telemetry", "t", false, "enable telemetry output during flight")
}

func enable(continuity, apogee, telemetry bool) {
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

	if continuity {
		serialport.SendCommand(port, "C")
		time.Sleep(5 * time.Second)
		serialport.SendCommand(port, "C")
	}

	if apogee {
		serialport.SendCommand(port, "A1")
	}
	if telemetry {
		serialport.SendCommand(port, "T1")
	}
}
