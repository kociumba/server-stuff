package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arl/statsviz"
)

// Initiate initiates the handlers and starts the server (also initiates a stats server on port 8080/debug/statsviz).
//
// This function does not take any parameters.
// It does not return any values.
func Initiate() {

	// =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=
	// =-  I N I T I A T E   T H E   H A N D L E R S  -=
	// =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

	http.HandleFunc("/", DomainHandler)

	statsviz.Register(http.DefaultServeMux, statsviz.TimeseriesPlot(scatterPlot()))

	// =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=
	// =-    I N I T I A T E   T H E   S E R V E R    -=
	// =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

	// server stats on port 8080
	// might change the stats port to be diffrent later
	fmt.Println("Point your browser to http://localhost:8080/debug/statsviz/")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

// myGraph generates a statsviz.TimeSeriesPlot.
//
// It defines a custom time series struct called myTimeSeries with the fields Name, Unitfmt, and GetValue.
// The GetValue field is a function that returns the value of the time series.
// The function then builds a new plot using the TimeSeriesPlotConfig struct, with the provided configuration.
// It returns the generated plot.
func scatterPlot() statsviz.TimeSeriesPlot {
	// Describe the 'balls' time series.

	balls := statsviz.TimeSeries{
		Name:    "balls",
		Unitfmt: "%{y:.4s}",
		GetValue: func() float64 {
			return checkBalls()
		},
	}

	// Build a new plot, showing our balls time series
	plot, err := statsviz.TimeSeriesPlotConfig{
		Name:       "balls",
		Title:      "Balls",
		Type:       statsviz.Scatter,
		InfoText:   `is balls ?`,
		YAxisTitle: "no balls/yes balls",
		Series:     []statsviz.TimeSeries{balls},
	}.Build()
	if err != nil {
		log.Fatalf("failed to build timeseries plot: %v", err)
	}

	return plot
}

var ballData = 0.0

// checkBalls returns the value of the data variable after checking the is_balls channel.
//
// It does not take any parameters.
// It returns a float64 value.
func checkBalls() float64 {

	select {
	case value := <-is_balls:
		if value {
			ballData = 1.0
		}
	default:
		break
	}

	return ballData
}
