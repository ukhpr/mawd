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
	"strconv"

	"github.com/spf13/cobra"
	"github.com/ukhpr/mawd/serialport"
	"go.bug.st/serial"
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Report altimeter statistics",
	Long:  "Report stats (ground, apogee, #samps, machdel, mainalt)",
	Run: func(cmd *cobra.Command, args []string) {
		stats()
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}

func stats() {
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

	// Request the altimeter stats information
	serialport.SendCommand(port, "S")

	// Read and print the response
	// expected respose is CRLF followed by 5 lines
	// each line is 1 letter followed by 5 digits followed by CRLF
	// G=ground, A=altitude, N=no. of samples, D=mach delay, M=main altitude
	line := serialport.Readline(port) // Read and discard initial blank line

	line = serialport.Readline(port) // Read 1st line of response (ground)
	if n, err := strconv.Atoi(line[1:6]); err == nil {
		fmt.Println("Detected Ground Altitude:", n, "ft")
	} else {
		fmt.Println(err)
	}

	line = serialport.Readline(port) // Read 2nd line of response (altitude)
	if n, err := strconv.Atoi(line[1:6]); err == nil {
		fmt.Println("Altitude at Apogee:", n, "ft AGL")
	} else {
		fmt.Println(err)
	}

	line = serialport.Readline(port) // Read 3rd line of response (number of samples)
	if n, err := strconv.Atoi(line[1:6]); err == nil {
		fmt.Println("Number of Samples:", n)
	} else {
		fmt.Println(err)
	}

	line = serialport.Readline(port) // Read 4th line of response (mach delay)
	if n, err := strconv.Atoi(line[1:6]); err == nil {
		fmt.Println("Mach Delay:", n, "seconds")
	} else {
		fmt.Println(err)
	}

	line = serialport.Readline(port) // Read 5th line of response (main altitude)
	if n, err := strconv.Atoi(line[1:6]); err == nil {
		fmt.Println("Main Deployment Altitude:", n, "ft AGL")
	} else {
		fmt.Println(err)
	}
}
