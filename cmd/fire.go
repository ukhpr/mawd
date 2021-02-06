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
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/ukhpr/mawd/serialport"
	"go.bug.st/serial"
)

// fireCmd represents the fire command
var fireCmd = &cobra.Command{
	Use:   "fire",
	Short: "Fire drogue or main pyro channel",
	Run: func(cmd *cobra.Command, args []string) {
		channel, _ := cmd.Flags().GetString("channel")
		fmt.Println("Pyro channel selected =", channel)
		fire(channel)
	},
}

func init() {
	rootCmd.AddCommand(fireCmd)

	fireCmd.Flags().StringP("channel", "c", "", "channel to fire")
	fireCmd.MarkFlagRequired("channel")
}

func fire(channel string) {
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

	fmt.Println("Sending pyro channel fire request...")

	// Request the altimeter to fire pyro channel
	switch channel {
	case "main":
		serialport.SendCommand(port, "FM")
	case "drogue":
		serialport.SendCommand(port, "FD")
	default:
		fmt.Println("Invalid pyro channel", channel)
	}
}
