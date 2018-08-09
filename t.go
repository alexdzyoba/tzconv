package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/sahilm/fuzzy"
	"github.com/spf13/pflag"
)

func usage() {
	fmt.Printf("Usage: %s [options] <target timezone> [source time] [source timezone]\n", os.Args[0])
	pflag.PrintDefaults()
}

func formatSelect(flag string) string {
	switch flag {
	case "u":
		fallthrough
	case "unix":
		return time.UnixDate

	case "1123":
		fallthrough
	case "rfc1123":
		fallthrough
	case "RFC1123":
		return time.RFC1123

	case "3339":
		fallthrough
	case "rfc3339":
		fallthrough
	case "RFC3339":
		return time.RFC3339

	case "t":
		fallthrough
	case "time":
		return "15:04"

	case "t2":
		fallthrough
	case "time2":
		fallthrough
	case "kitchen":
		return time.Kitchen

	default:
		return flag
	}
}

func printLocations() {
	sort.Strings(Locations)
	for _, l := range Locations {
		fmt.Println(l)
	}
}

func loadLocation(s string) (*time.Location, error) {
	resultSet := fuzzy.Find(s, Locations)

	if len(resultSet) == 0 {
		return nil, fmt.Errorf("error: unable to find timezone '%s'", s)
	}

	resultIndex := resultSet[0].Index
	location, err := time.LoadLocation(Timezones[resultIndex])
	if err != nil {
		return nil, err
	}

	return location, nil
}

func tzconv(argc int) (*time.Time, error) {
	sourceTime := time.Now()
	sourceLocation := time.Local
	targetLocation := time.Local

	var err error
	switch argc {
	case 3: //tzconv <target timezone> <source time> <source timezone>
		sourceLocation, err = loadLocation(pflag.Arg(2))
		if err != nil {
			return nil, err
		}
		fallthrough // to handle source time and target timezone

	case 2: //tzconv <target tz> <source time>
		parsedTime, err := time.Parse("15:04", pflag.Arg(1))
		if err != nil {
			return nil, err
		}

		// add date information to the parsed time
		now := time.Now()
		sourceTime = time.Date(now.Year(), now.Month(), now.Day(), parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second(), 0, sourceLocation)
		fallthrough // to handle target timezone

	case 1: // tzconv <target tz>
		targetLocation, err = loadLocation(pflag.Arg(0))
		if err != nil {
			return nil, err
		}

	case 0: // tzconv
		// just print the current time

	default:
		usage()
		return nil, errors.New("error: invalid arguments")
	}

	targetTime := sourceTime.In(targetLocation)
	return &targetTime, nil
}

func main() {
	pflag.Usage = usage
	pflag.ErrHelp = errors.New("error: help requested")

	formatFlag := pflag.StringP("format", "f", "time", "Time format (time, unix, rfc1123, rfc3339, kitchen)")
	printFlag := pflag.BoolP("print", "p", false, "Print available locations")
	pflag.Parse()

	if *printFlag {
		printLocations()
		os.Exit(0)
	}

	argc := len(pflag.Args())

	targetTime, err := tzconv(argc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	format := formatSelect(*formatFlag)
	fmt.Println(targetTime.Format(format))
}
