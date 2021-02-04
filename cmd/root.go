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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ukhpr/mawd/serialport"
)

var (
	cfgFile  string
	port     string
	baudRate int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mawd",
	Short: "MAWD Altimeter Data Capture",
	Long:  `A tool to configure a PerfectFlite MAWD altimeter and download flight data.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "configuration file")
	rootCmd.PersistentFlags().StringVar(&port, "port", "", "serial port")
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	rootCmd.PersistentFlags().IntVar(&baudRate, "baudrate", 38400, "baud rate")
	viper.BindPFlag("baudRate", rootCmd.PersistentFlags().Lookup("baudrate"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find user config directory.
		config, err := os.UserConfigDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".mawd" (without extension).
		viper.AddConfigPath(config + string(os.PathSeparator) + "mawd")
		viper.SetConfigName("mawd")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	port = viper.GetString("port")
	fmt.Println("Using serial port:", port)
	err := serialport.IsValid(port)
	if err != nil {
		fmt.Println("Invalid Serial Port:", err)
		os.Exit(1)
	}

	baudRate = viper.GetInt("baudrate")
	fmt.Println("Using baud rate:", baudRate)

	fmt.Println("")
}
