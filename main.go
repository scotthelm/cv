package main

import (
	"flag"
	"fmt"
	"sort"
)

func main() {
	flag.Parse()
	if *showUnits {
		listUnits()
	} else {
		performConversion()
	}
}

var num = flag.Float64("n", 1, "number to convert")
var uom = flag.String("i", "f", "input unit of measure")
var out = flag.String("o", "c", "output unit of measure")
var showUnits = flag.Bool("u", false, "show accepted units")

func listUnits() {
	fmt.Println("Accepted Units")
	sortable := make([]string, len(acceptedUnits))
	counter := 0
	for k, v := range acceptedUnits {
		sortable[counter] = fmt.Sprintf("%s (%v)", k, v)
		counter++
	}
	sort.Strings(sortable)
	for _, v := range sortable {
		fmt.Println(v)
	}
}

func performConversion() {
	fn := conversions[*uom][*out]
	if fn != nil {
		fmt.Println(fmtCnv(*num, *uom, *out, fn(*num)))
	} else {
		fmt.Println("Conversion Not Found")
	}
}

type conversion func(in float64) float64

func fmtCnv(input float64, uom string, out string, converted float64) string {
	return fmt.Sprintf(
		"%.1f %s = %.1f %s",
		input,
		unit(uom),
		converted,
		unit(out),
	)
}

func unit(u string) string {
	return acceptedUnits[u]
}

var acceptedUnits = map[string]string{
	"f":  "Fahrenheit",
	"c":  "Celsius",
	"km": "Kilometers",
	"mi": "Miles",
	"m":  "Meters",
	"Z":  "Zach",
	"J":  "Jamey",
}

var self conversion = func(in float64) float64 { return in * 1.0 }

var conversions = map[string]map[string]conversion{
	"f": {
		"f": self,
		"c": func(in float64) float64 { return (in - 32.0) * 5.0 / 9.0 },
		"k": func(in float64) float64 { return ((in - 32.0) * 5.0 / 9.0) - 273.15 },
	},
	"c": {
		"c": self,
		"f": func(in float64) float64 { return in*9.0/5.0 + 32.0 },
		"k": func(in float64) float64 { return in + 273.15 },
	},
	"km": {
		"mi": func(in float64) float64 { return in * 0.621371 },
		"km": self,
		"m":  func(in float64) float64 { return in * 1000 },
	},
	"mi": {
		"mi": self,
		"km": func(in float64) float64 { return in * 1.609343502101154 },
		"m":  func(in float64) float64 { return in * 1.609343502101154 * 1000 },
	},
	"Z": {
		"Z": self,
		"J": func(in float64) float64 { return in / 0.64197530864 },
	},
	"J": {
		"J": self,
		"Z": func(in float64) float64 { return in * 0.64197530864 },
	},
}
