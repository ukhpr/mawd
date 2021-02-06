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

	"github.com/spf13/cobra"
	"github.com/ukhpr/mawd/serialport"
	"go.bug.st/serial"
)

// disableCmd represents the disable command
var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable function on altimeter",
	Long:  `Disable either apogee delay or telemetry output`,
	Run: func(cmd *cobra.Command, args []string) {
		apogee, _ := cmd.Flags().GetBool("apogee")
		telemetry, _ := cmd.Flags().GetBool("telemetry")
		disable(apogee, telemetry)
	},
}

func init() {
	rootCmd.AddCommand(disableCmd)

	disableCmd.Flags().BoolP("apogee", "a", false, "disable 1 second apogee delay")
	disableCmd.Flags().BoolP("telemetry", "t", false, "disable telemetry output during flight")
}

func disable(apogee, telemetry bool) {
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

	if apogee {
		serialport.SendCommand(port, "A0")
	}
	if telemetry {
		serialport.SendCommand(port, "T0")
	}
}
