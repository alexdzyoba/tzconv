# tzconv

[![Go Report Card](https://goreportcard.com/badge/github.com/alexdzyoba/tzconv)](https://goreportcard.com/report/github.com/alexdzyoba/tzconv)
[![CircleCI](https://circleci.com/gh/alexdzyoba/tzconv.svg?style=svg)](https://circleci.com/gh/alexdzyoba/tzconv)

Time converter with timezone fuzzy search

## Examples

    $ tzconv -h
    Usage: tzconv [options] <target timezone> [source time] [source timezone]
      -f, --format string   Time format (time, unix, rfc1123, rfc3339, kitchen) (default "time")
      -p, --print           Print available locations

By default, `tzconv` shows current time in current timezone in simple format

    $ tzconv
    15:58

If you pass a timezone it will adjust current time to the supplied timezone

    $ tzconv toronto
    11:58

To get a list of timezones, just invoke it with `-p` option

    $ tzconv -p
    Abidjan
    Accra
    Adak
    Adelaide
    Algiers
    Almaty
    Amman
    Amsterdam
    Anadyr
    Anchorage
    ...

Timezones are taken from the [tz database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)
but takes the location name (without continent or ocean) and fuzzy searched so
you can just type

    $ tzconv trnt
    12:06

Another form of `tzconv` invocation is to show specific time in a given timezone

    $ tzconv toronto 15:00
    11:00


And final form is to convert time between timezones, e.g. this will show 15:00
of Moscow time in Toronto:

    $ tzconv toronto 15:00 moscow
    08:00

And of course you can type partial names of timezones for fuzzy search

    $ tzconv trnt 15:00 mscw
    08:00

There are also different time formats to output the final time like unix

    $ tzconv -f unix toronto
    Tue Aug 14 11:59:17 EDT 2018

or RFC3339

    $ tzconv -f rfc3339 toronto
    2018-08-14T11:59:35-04:00

RFC format names may be supplied with digits only:

    $ tzconv -f 3339 toronto
    2018-08-14T12:00:08-04:00

## Install

Grab the latest prebuilt release from Github
https://github.com/alexdzyoba/tzconv/releases/latest

There are archives with statically linked binaries for Linux and macOS as well
as deb and RPM packages for Linux.

As an alternative you can install it with `go install` if `$GOBIN` is in your
`$PATH`:

    $ go install github.com/alexdzyoba/tzconv@latest
    $ tzconv
    16:22

