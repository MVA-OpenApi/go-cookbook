/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	printer "printer/src"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "printer",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var seperator string

var namesCmd = &cobra.Command{
	Use:   "names",
	Short: "Print names from .json file",
	Long: `Print names from .json file
For example: printer names haha.json`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		if printer.CheckIfFileExists(path) != true {
			return
		}
		jsonFile, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var names printer.Names

		json.Unmarshal(byteValue, &names)

		for _, name := range names.Names {
			fmt.Print(name)
			fmt.Print(seperator)
		}
		fmt.Println()

	},
}

var citiesCmd = &cobra.Command{
	Use:   "cities",
	Short: "Print cities from .json file",
	Long: `Print cities from .json file
For example: printer cities haha.json`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		if printer.CheckIfFileExists(path) != true {
			return
		}
		jsonFile, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var cities printer.Cities

		json.Unmarshal(byteValue, &cities)

		for _, city := range cities.Cities {
			fmt.Print(city)
			fmt.Print(seperator)
		}
		fmt.Println()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.example.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// add seperator flag
	rootCmd.PersistentFlags().StringVarP(&seperator, "seperator", "s", " ", "Seperator between entries when printing")

	// add sub commands
	rootCmd.AddCommand(namesCmd)
	rootCmd.AddCommand(citiesCmd)
}
