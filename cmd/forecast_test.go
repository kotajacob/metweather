package cmd

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestForecast(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/localForecastDunedin", func(w http.ResponseWriter, r *http.Request) {
		reader, err := os.Open("data/localForecastDunedin.json")
		if err != nil {
			t.Errorf("metweather.forecast returned error: %v", err)
		}
		io.Copy(w, reader)
	})

	ctx := context.Background()
	out = new(bytes.Buffer) // captured output
	forecastWeekly(client, ctx)
	got := out.(*bytes.Buffer).String()
	want := `2021-7-16 Partly cloudy 7-13c
2021-7-17 Showers 8-11c
2021-7-18 Few showers 7-10c
2021-7-19 Showers 4-11c
2021-7-20 Partly cloudy 6-11c
2021-7-21 Partly cloudy 6-12c
2021-7-22 Few showers 6-11c
2021-7-23 Fine 5-10c
2021-7-24 Fine 5-11c
2021-7-25 Partly cloudy 6-12c
`
	if got != want {
		t.Errorf("got = %q\nwant = %q", got, want)
	}
}
