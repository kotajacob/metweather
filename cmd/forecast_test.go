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
	want := `2021-7-16 Partly cloudy 7-13°C
2021-7-17 Showers 8-11°C
2021-7-18 Few showers 7-10°C
2021-7-19 Showers 4-11°C
2021-7-20 Partly cloudy 6-11°C
2021-7-21 Partly cloudy 6-12°C
2021-7-22 Few showers 6-11°C
2021-7-23 Fine 5-10°C
2021-7-24 Fine 5-11°C
2021-7-25 Partly cloudy 6-12°C
`
	if got != want {
		t.Errorf("metweather.forecastWeekly\ngot = %q\nwant = %q", got, want)
	}
}
