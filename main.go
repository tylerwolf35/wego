package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/schachmat/ingo"
	_ "github.com/tylerwolf35/wego/backends"
	_ "github.com/tylerwolf35/wego/frontends"
	"github.com/tylerwolf35/wego/iface"
)

func pluginLists() {
	bEnds := make([]string, 0, len(iface.AllBackends))
	for name := range iface.AllBackends {
		bEnds = append(bEnds, name)
	}
	sort.Strings(bEnds)

	fEnds := make([]string, 0, len(iface.AllFrontends))
	for name := range iface.AllFrontends {
		fEnds = append(fEnds, name)
	}
	sort.Strings(fEnds)

	fmt.Fprintln(os.Stderr, "Available backends:", strings.Join(bEnds, ", "))
	fmt.Fprintln(os.Stderr, "Available frontends:", strings.Join(fEnds, ", "))
}

func main() {
	// initialize backends and frontends (flags and default config)
	for _, be := range iface.AllBackends {
		be.Setup()
	}
	for _, fe := range iface.AllFrontends {
		fe.Setup()
	}

	// initialize global flags and default config
	location := flag.String("location", "37.4275,-122.1697", "`LOCATION` to be queried")
	flag.StringVar(location, "l", "37.4275,-122.1697", "`LOCATION` to be queried (shorthand)")
	numdays := flag.Int("days", 3, "`NUMBER` of days of weather forecast to be displayed")
	flag.IntVar(numdays, "d", 3, "`NUMBER` of days of weather forecast to be displayed (shorthand)")
	unitSystem := flag.String("units", "metric", "`UNITSYSTEM` to use for output.\n    \tChoices are: metric, imperial, si, metric-ms")
	flag.StringVar(unitSystem, "u", "metric", "`UNITSYSTEM` to use for output. (shorthand)\n    \tChoices are: metric, imperial, si, metric-ms")
	selectedBackend := flag.String("backend", "darksky", "`BACKEND` to be used")
	flag.StringVar(selectedBackend, "b", "darksky", "`BACKEND` to be used (shorthand)")
	selectedFrontend := flag.String("frontend", "ascii-art-table", "`FRONTEND` to be used")
	flag.StringVar(selectedFrontend, "f", "ascii-art-table", "`FRONTEND` to be used (shorthand)")

	// print out a list of all backends and frontends in the usage
	tmpUsage := flag.Usage
	flag.Usage = func() {
		tmpUsage()
		pluginLists()
	}

	// read/write config and parse flags
	if err := ingo.Parse("wego"); err != nil {
		log.Fatalf("Error parsing config: %v", err)
	}

	// non-flag shortcut arguments overwrite possible flag arguments
	for _, arg := range flag.Args() {
		if v, err := strconv.Atoi(arg); err == nil && len(arg) == 1 {
			*numdays = v
		} else {
			*location = arg
		}
	}

	// get selected backend and fetch the weather data from it
	be, ok := iface.AllBackends[*selectedBackend]
	if !ok {
		log.Fatalf("Could not find selected backend \"%s\"", *selectedBackend)
	}
	r := be.Fetch(*location, *numdays)

	// set unit system
	unit := iface.UnitsMetric
	if *unitSystem == "imperial" {
		unit = iface.UnitsImperial
	} else if *unitSystem == "si" {
		unit = iface.UnitsSi
	} else if *unitSystem == "metric-ms" {
		unit = iface.UnitsMetricMs
	}

	// get selected frontend and render the weather data with it
	fe, ok := iface.AllFrontends[*selectedFrontend]
	if !ok {
		log.Fatalf("Could not find selected frontend \"%s\"", *selectedFrontend)
	}
	fe.Render(r, unit)
}
