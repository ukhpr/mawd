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
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/ukhpr/mawd/serialport"
	"go.bug.st/serial"
)

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump flight data from the altimeter",
	Long: `Dump the most recent flight data from the altimeter.
A filename can optionally be provided.`,
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := cmd.Flags().GetString("file")
		dump(filename)
	},
}

func init() {
	rootCmd.AddCommand(dumpCmd)

	dumpCmd.Flags().StringP("file", "f", "mawddump.txt", "Output file name")
}

func dump(filename string) {
	fmt.Println("Output filename:", filename)

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

	// Get the number of samples of recorded data
	sampleCount := getSampleCount(port)
	fmt.Println("Sample Count:", sampleCount)

	// Dump the recorded data
	serialport.SendCommand(port, "D")

	// Process response
	line := serialport.Readline(port) // Read blank line
	line = serialport.Readline(port)  // Read 1st line of response (sample rate info)
	fmt.Println(line)

	// Extract sample interval from header
	sampleInterval, err := strconv.Atoi(line[20:22])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Sample Interval:", sampleInterval, "ms")

	line = serialport.Readline(port) // Read blank line

	var data [][]string

	for i := 0; i < sampleCount; i++ {
		line = serialport.Readline(port) // Read next line of response
		n, err := strconv.Atoi(line[1:5])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(i, ",", i*sampleInterval, ",", n)
		data[i][0] = strconv.Itoa(i * sampleInterval)
		data[i][1] = strconv.Itoa(n)
	}

	// Write data to file
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			log.Fatal("Cannot write to file", err)
		}
	}
}

func getSampleCount(port serial.Port) int {
	// sampleCount := 0
	// Request the altimeter stats information
	serialport.SendCommand(port, "S")

	// Read and print the response
	// expected respose is CRLF followed by 5 lines
	// each line is 1 letter followed by 5 digits followed by CRLF
	// G=ground, A=altitude, N=no. of samples, D=mach delay, M=main altitude
	line := serialport.Readline(port) // Read and discard initial blank line
	line = serialport.Readline(port)  // Read 1st line of response (ground)
	line = serialport.Readline(port)  // Read 2nd line of response (altitude)
	line = serialport.Readline(port)  // Read 3rd line of response (number of samples)
	sampleCount, err := strconv.Atoi(line[1:6])
	if err != nil {
		fmt.Println(err)
	}
	// if sampleCount, err := strconv.Atoi(line[1:6]); err == nil {
	// 	fmt.Println("Number of Samples:", sampleCount)
	// } else {
	// 	fmt.Println(err)
	// }

	line = serialport.Readline(port) // Read 4th line of response (mach delay)
	line = serialport.Readline(port) // Read 5th line of response (main altitude)

	return sampleCount
}
