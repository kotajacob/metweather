package cmd

import (
	"context"
	"fmt"
	"log"
	"math"

	"git.sr.ht/~kota/metservice-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// observationCmd represents the observation command
var observationCmd = &cobra.Command{
	Use:   "observation",
	Short: "display current or past weather observations",
	Run:   observation,
}

func init() {
	rootCmd.AddCommand(observationCmd)
	observationCmd.PersistentFlags().StringP("location", "l", "", "Location used for the weather observation")
	viper.BindPFlag("location", observationCmd.PersistentFlags().Lookup("location"))
}

// observation fetches and prints an observation based on options provided
func observation(cmd *cobra.Command, args []string) {
	client := metservice.NewClient()
	ctx := context.Background()
	location := viper.GetString("location")
	if location == "" {
		log.Fatal("location is required either using the flag or config")
	}
	err := observationThreeHour(client, ctx, location)
	if err != nil {
		log.Fatal(err)
	}
}

// observationThreeHour fetches and prints an observation of the last three
// hours
func observationThreeHour(client *metservice.Client, ctx context.Context, location string) error {
	o, _, err := client.GetObservation(ctx, location)
	if err != nil {
		return fmt.Errorf("getting observation: %v", err)
	}
	fmt.Fprintf(out, "%d°C ", *o.ThreeHour.Temp)
	fmt.Fprintf(out, "%dkm/h ", *o.ThreeHour.WindSpeed)
	fmt.Fprintf(out, "%s - ", *o.ThreeHour.WindDirection)
	fmt.Fprintf(out, "%d%%", *o.ThreeHour.Humidity)
	fmt.Fprintf(out, "(%.2f°C)\n", feelsLike(*o.ThreeHour.Temp, *o.ThreeHour.Humidity, *o.ThreeHour.WindSpeed))
	return nil
}

// feelsLike calculates an estimated temperate that considers your ability to
// heat or cool yourself. We use the same formula described by Metservice here
// https://blog.metservice.com/FeelsLikeTemp
// When the measured air temperature is below 10°C the windChill algorithm is
// used. If it is above 14°C the apparentTemp is calculated and used if it's it
// higher than the measured air temp. For values between 10-14°C a pragmatic
// linear roll-off of the wind chill is used. So 12°C at 5km/h wind has a
// windChill of 12°C.
// NOTE: The windChill formula listen on the Metservice website is incorrect.
func feelsLike(temp int, humidity int, speed int) float64 {
	// return whichever is higher of apparentTemp and temp
	if temp >= 14 {
		t := float64(temp)
		at := apparentTemp(temp, humidity, speed)
		return math.Max(t, at)
	}
	w := windChill(temp, speed)
	// return windChill
	if temp < 10 {
		return w
	}
	// linear roll-off between windChill and temp
	t := float64(temp)
	f := t - (((t - w) * (14 - t)) / 4)
	return f
}

// windChill calculates the wind chill temperature in degrees Celsius based on
// the following algorithm
// https://web.archive.org/web/20060415000715/http://www.msc.ec.gc.ca/education/windchill/Science_equations_e.cfm
// w = 13.12 + 0.6215*t - 11.35*k^0.16 + 0.396*t*k^0.16
// t = Dry bulb temperature (°C)
// k = Average wind speed in km/h
func windChill(temp int, speed int) float64 {
	t := float64(temp)
	k := math.Pow(float64(speed), 0.16)
	w := 13.12 + 0.6215*t - 11.35*k + 0.396*t*k
	return w
}

// apparentTemp calculates the apparent temperature in degrees Celsius based on
// the following algorithm http://www.bom.gov.au/info/thermal_stress/
// at = t + 0.33e - 0.70m - 4.00
// t = Dry bulb temperature (°C)
// e = Water vapour pressure (hPa) [humidity]
// m = Wind speed (m/s) at an elevation of 10 meters
// our function takes the values in the following formats and converts them
// temp = Air temperature (°C) int
// humidity = Relative humidity percentage int (0-100)
// speed = Wind speed km/h int
func apparentTemp(temp int, humidity int, speed int) float64 {
	t := float64(temp)
	e := float64(humidity) / 100 * 6.105 * (17.27 * t / (237.7 + t))
	m := float64(speed) * (5 / 18)
	at := t + 0.33*e - 0.70*m - 4.00
	return at
}
