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
	Use:   "printer [property] [input file path] [output file path] [flags]",
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
	Use:   "names [input file path] [output file path] [flags]",
	Short: "Create a main.go file that can print the names.",
	Long: `Create a main.go file that can print the names from the input JSON-File.
For example: printer names ./input.json ./ouput/`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		input_path := args[0]
		output_path := args[1]
		if !printer.CheckIfFileExists(input_path) {
			fmt.Println("No valid input file path given.")
			return
		}
		jsonFile, err := os.Open(input_path)
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var names printer.Names

		json.Unmarshal(byteValue, &names)

		printer.GenerateGoFile(names.Names, output_path)
	},
}

var citiesCmd = &cobra.Command{
	Use:   "cities [input file path] [output file path] [flags]",
	Short: "Create a main.go file that can print the names.",
	Long: `Create a main.go file that can print the names from the input JSON-File.
For example: printer cities ./input.json ./ouput/`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		input_path := args[0]
		output_path := args[1]
		if !printer.CheckIfFileExists(input_path) {
			return
		}
		jsonFile, err := os.Open(input_path)
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var cities printer.Cities

		json.Unmarshal(byteValue, &cities)

		printer.GenerateGoFile(cities.Cities, output_path)
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
