/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/cwpurdy/gander/csv"
	"strconv"

	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
)

// shapeCmd represents the shape command
var shapeCmd = &cobra.Command{
	Use:   "shape",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		shape, err := csv.GetShape(args[0], headless)

		if err != nil {
			fmt.Println("ERR", err)
		} else {

			p := termenv.ColorProfile()

			rows := termenv.String(strconv.Itoa(shape.Rows)).Background(p.Color("#000000")).Foreground(p.Color("#FFFFFF"))
			cols := termenv.String(strconv.Itoa(shape.Cols)).Background(p.Color("#000000")).Foreground(p.Color("#FFFFFF"))

			Rtitle := termenv.String(" ROWS: ").Background(p.Color("#FF00FF")).Foreground(p.Color("#000000")).Bold()
			Ctitle := termenv.String(" COLS: ").Background(p.Color("#FF00FF")).Foreground(p.Color("#000000")).Bold()

			fmt.Println(Rtitle, rows," | ", Ctitle, cols)
		}
	},
}

func init() {
	rootCmd.AddCommand(shapeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shapeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shapeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
