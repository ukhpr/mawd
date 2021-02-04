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
	"go.bug.st/serial/enumerator"
)

// portsCmd represents the ports command
var portsCmd = &cobra.Command{
	Use:   "ports",
	Short: "List available serial ports",
	Run: func(cmd *cobra.Command, args []string) {
		listports()
	},
}

func init() {
	rootCmd.AddCommand(portsCmd)
}

func listports() {
	// ports, err := serial.GetPortsList()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if len(ports) == 0 {
	// 	log.Fatal("No serial ports found!")
	// }
	// for _, port := range ports {
	// 	fmt.Printf("Found port: %v\n", port)
	// }

	usbports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(usbports) == 0 {
		fmt.Println("No serial ports found!")
		return
	}
	for _, port := range usbports {
		fmt.Printf("Found port: %s\n", port.Name)
		if port.IsUSB {
			fmt.Printf("   USB ID     %s:%s\n", port.VID, port.PID)
			fmt.Printf("   USB serial %s\n", port.SerialNumber)
		}
	}
}
