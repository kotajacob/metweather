/*
Copyright © 2021 Dakota Walsh <kota@nilsu.org>

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
	"context"
	"fmt"

	"git.sr.ht/~kota/metservice-go"
	"github.com/spf13/cobra"
)

// forecastCmd represents the forecast command
var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: Forecast,
}

func init() {
	rootCmd.AddCommand(forecastCmd)
}

func Forecast(cmd *cobra.Command, args []string) {
	client := metservice.NewClient()
	ctx := context.Background()

	forecast, _, err := client.GetForecast(ctx, "Dunedin")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*forecast.LocationIPS)
	for _, day := range forecast.Days {
		fmt.Printf("%v\nforecast: %v\nmax: %vC\nmin: %vC\n\n",
			*day.Date,
			*day.ForecastWord,
			*day.Max,
			*day.Min)
	}
}
