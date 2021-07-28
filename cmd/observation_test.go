package cmd

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestObservationThreeHour(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/localObs_Dunedin", func(w http.ResponseWriter, r *http.Request) {
		reader, err := os.Open("data/localObs_Dunedin.json")
		if err != nil {
			t.Errorf("metweather.observationThreeHour returned error: %v", err)
		}
		io.Copy(w, reader)
	})

	ctx := context.Background()
	out = new(bytes.Buffer) // captured output
	err := observationThreeHour(client, ctx, "Dunedin")
	if err != nil {
		t.Errorf("metweather.observationThreeHour returned error: %v", err)
	}
	got := out.(*bytes.Buffer).String()
	want := "9°C 2km/h W - 77%(10.01°C)\n"
	if got != want {
		t.Errorf("metweather.observationThreeHour\ngot = %q\nwant = %q", got, want)
	}
}

func TestWindChill(t *testing.T) {
	var tests = []struct {
		temp  int
		speed int
		want  float64
	}{
		{12, 5, 12.042135509662362},
		{12, 10, 11.04098839261837},
		{12, 25, 9.535110852166257},
		{12, 50, 8.239921987968058},
		{9, 5, 8.640714167661585},
		{9, 10, 7.459305944972208},
		{9, 25, 5.682288283565697},
		{9, 50, 4.153894755731934},
		{0, 5, -1.5635498583407408},
		{0, 10, -3.2857413979662784},
		{0, 25, -5.87617942223598},
		{0, 50, -8.104186940976438},
	}
	for _, test := range tests {
		if got := windChill(test.temp, test.speed); got != test.want {
			t.Errorf("metweather.windChill\ntemp = %d\nspeed = %d\nwant = %f\n got = %f\n",
				test.temp,
				test.speed,
				test.want,
				got)
		}
	}
}
