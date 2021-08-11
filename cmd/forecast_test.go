package cmd

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestForecastWeekly(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/localForecastDunedin", func(w http.ResponseWriter, r *http.Request) {
		reader, err := os.Open("data/localForecastDunedin.json")
		if err != nil {
			t.Errorf("metweather.forecastWeekly returned error: %v", err)
		}
		io.Copy(w, reader)
	})

	ctx := context.Background()
	out = new(bytes.Buffer) // captured output
	err := forecastWeekly(client, ctx, "Dunedin")
	if err != nil {
		t.Errorf("metweather.forecastWeekly returned error: %v", err)
	}
	got := out.(*bytes.Buffer).String()
	want := `2021-7-16 7-13°C Partly cloudy
2021-7-17 8-11°C Showers
2021-7-18 7-10°C Few showers
2021-7-19 4-11°C Showers
2021-7-20 6-11°C Partly cloudy
2021-7-21 6-12°C Partly cloudy
2021-7-22 6-11°C Few showers
2021-7-23 5-10°C Fine
2021-7-24 5-11°C Fine
2021-7-25 6-12°C Partly cloudy
`
	if got != want {
		t.Errorf("metweather.forecastWeekly\ngot = %q\nwant = %q", got, want)
	}
}
