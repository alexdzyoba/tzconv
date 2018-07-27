package main

import (
	"fmt"
	"os"
	"time"
)

func usage() {
	fmt.Printf("Usage: %s <target timezone> [source time] [source timezone]\n", os.Args[0])
}

func main() {
	argc := len(os.Args)

	sourceTime := time.Now()
	sourceLocation := time.Local
	targetLocation := time.Local

	var err error
	switch argc {
	case 4: //tzconv <target timezone> <source time> <source timezone>
		sourceLocation, err = time.LoadLocation(os.Args[3])
		if err != nil {
			panic(err)
		}
		fallthrough // to handle source time and target timezone

	case 3: //tzconv <target tz> <source time>
		parsedTime, err := time.Parse("15:04", os.Args[2])
		if err != nil {
			panic(err)
		}

		// add date information to the parsed time
		now := time.Now()
		sourceTime = time.Date(now.Year(), now.Month(), now.Day(), parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second(), 0, sourceLocation)
		fallthrough // to handle target timezone

	case 2: // tzconv <target tz>
		targetLocation, err = time.LoadLocation(os.Args[1])
		if err != nil {
			panic(err)
		}

	case 1: // tzconv
		// just print the current time

	default:
		usage()
		os.Exit(1)
	}

	fmt.Println(sourceTime.In(targetLocation))
}
