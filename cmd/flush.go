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
	"go.bug.st/serial"
)

// flushCmd represents the flush command
var flushCmd = &cobra.Command{
	Use:   "flush",
	Short: "Flush the Serial Port buffers",
	Long:  "Flush the Serial Port buffers",
	Run: func(cmd *cobra.Command, args []string) {
		flush()
	},
}

func init() {
	rootCmd.AddCommand(flushCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// flushCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// flushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func flush() {
	// Open serial port
	mode := &serial.Mode{
		BaudRate: 38400,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	p, err := serial.Open(port, mode)
	if err != nil {
		log.Fatal(err)
	}

	err = p.ResetInputBuffer()
	if err != nil {
		fmt.Println(err)
	}
	err = p.ResetOutputBuffer()
	if err != nil {
		fmt.Println(err)
	}
}
