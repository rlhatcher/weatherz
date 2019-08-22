package cmd

/*
Copyright © 2019 Ronald Hatcher <ronaldhatcher@mac.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

import (
	"fmt"
	"html/template"
	"os"

	"github.com/spf13/cobra"
)

// Summary structure for 3 day forecast
type Summary struct {
	Forecasts []Forecast
}

// Forecast structure for a single day
type Forecast struct {
	Icon        string
	High        string
	Low         string
	Description string
	Date        string
}

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate a 1-page ovrview of the forecast",
	Long: `Generates a 1-Page overview of the forecast intended as a hard copy.

Ideal for a handout.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("report called with supplier: " + supplier)
		tmpl := template.Must(template.ParseFiles("html/index.html"))
		forecast := metOffice()
		f, err := os.Create("html/report.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		tmpl.Execute(f, forecast)
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
