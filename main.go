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
	"f":   "Fahrenheit",
	"c":   "Celsius",
	"km":  "Kilometers",
	"mi":  "Miles",
	"m":   "Meters",
	"skm": "Square Kilometers",
	"smi": "Square Miles",
	"sm":  "Square Meters",
	"sy":  "Square Yards",
	"sf":  "Square Feet",
	"h":   "Hectares",
	"a":   "Acres",
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
	"skm": {
		"skm": self,
		"smi": func(in float64) float64 { return in * 0.386102 },
		"sm":  func(in float64) float64 { return in * 1e+6 },
		"sy":  func(in float64) float64 { return in * 1.196e+6 },
		"sf":  func(in float64) float64 { return in * 1.076e+7 },
		"h":   func(in float64) float64 { return in * 100 },
		"a":   func(in float64) float64 { return in * 247.105 },
	},
	"smi": {
		"skm": func(in float64) float64 { return in * 2.58999 },
		"smi": self,
		"sm":  func(in float64) float64 { return in * 2589990.001027 },
		"sy":  func(in float64) float64 { return in * 3.098e+6 },
		"sf":  func(in float64) float64 { return in * 2.788e+7 },
		"h":   func(in float64) float64 { return in * 259 },
		"a":   func(in float64) float64 { return in * 640 },
	},
	"sm": {
		"skm": func(in float64) float64 { return in * 1e-6 },
		"smi": func(in float64) float64 { return in * 3.861e-7 },
		"sm":  self,
		"sy":  func(in float64) float64 { return in * 1.196 },
		"sf":  func(in float64) float64 { return in * 10.7639 },
		"h":   func(in float64) float64 { return in * 1e-4 },
		"a":   func(in float64) float64 { return in * 0.000247105 },
	},
	"sy": {
		"skm": func(in float64) float64 { return in * 8.3613e-7 },
		"smi": func(in float64) float64 { return in * 3.2283e-7 },
		"sm":  func(in float64) float64 { return in * 0.8361300021625 },
		"sy":  self,
		"sf":  func(in float64) float64 { return in * 9 },
		"h":   func(in float64) float64 { return in * 8.3613e-5 },
		"a":   func(in float64) float64 { return in * 0.000206612 },
	},
	"sf": {
		"skm": func(in float64) float64 { return in * 9.2903e-8 },
		"smi": func(in float64) float64 { return in * 3.587e-8 },
		"sm":  func(in float64) float64 { return in * 0.092903 },
		"sy":  func(in float64) float64 { return in * 0.111111 },
		"sf":  self,
		"h":   func(in float64) float64 { return in * 9.2903e-6 },
		"a":   func(in float64) float64 { return in * 2.2957e-5 },
	},
	"h": {
		"skm": func(in float64) float64 { return in * 0.01 },
		"smi": func(in float64) float64 { return in * 0.00386102 },
		"sm":  func(in float64) float64 { return in * 10000 },
		"sy":  func(in float64) float64 { return in * 11959.9 },
		"sf":  func(in float64) float64 { return in * 107639 },
		"h":   self,
		"a":   func(in float64) float64 { return in * 2.47105 },
	},
	"a": {
		"skm": func(in float64) float64 { return in * 0.00404686 },
		"smi": func(in float64) float64 { return in * 0.0015625 },
		"sm":  func(in float64) float64 { return in * 4046.86 },
		"sy":  func(in float64) float64 { return in * 4840 },
		"sf":  func(in float64) float64 { return in * 43560 },
		"h":   func(in float64) float64 { return in * 0.404686 },
		"a":   self,
	},
}
