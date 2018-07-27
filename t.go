package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/pflag"
)

func usage() {
	fmt.Printf("Usage: %s <target timezone> [source time] [source timezone]\n", os.Args[0])
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

func main() {
	formatFlag := pflag.StringP("format", "f", "time", "Time format (time, unix, rfc1123, rfc3339, kitchen)")
	pflag.Parse()

	argc := len(pflag.Args())

	sourceTime := time.Now()
	sourceLocation := time.Local
	targetLocation := time.Local

	var err error
	switch argc {
	case 3: //tzconv <target timezone> <source time> <source timezone>
		sourceLocation, err = time.LoadLocation(pflag.Arg(2))
		if err != nil {
			panic(err)
		}
		fallthrough // to handle source time and target timezone

	case 2: //tzconv <target tz> <source time>
		parsedTime, err := time.Parse("15:04", pflag.Arg(1))
		if err != nil {
			panic(err)
		}

		// add date information to the parsed time
		now := time.Now()
		sourceTime = time.Date(now.Year(), now.Month(), now.Day(), parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second(), 0, sourceLocation)
		fallthrough // to handle target timezone

	case 1: // tzconv <target tz>
		targetLocation, err = time.LoadLocation(pflag.Arg(0))
		if err != nil {
			panic(err)
		}

	case 0: // tzconv
		// just print the current time

	default:
		usage()
		os.Exit(1)
	}

	targetTime := sourceTime.In(targetLocation)

	format := formatSelect(*formatFlag)
	fmt.Println(targetTime.Format(format))
}
