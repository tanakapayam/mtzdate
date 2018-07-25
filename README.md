# mtzdate [![Build Status](https://travis-ci.org/tanakapayam/mtzdate.svg?branch=master)](https://travis-ci.org/tanakapayam/mtzdate) [![Go Report Card](https://goreportcard.com/badge/github.com/tanakapayam/mtzdate)](https://goreportcard.com/report/github.com/tanakapayam/mtzdate)

*a command-line utility to show date in multiple time zones based on environment variables*

## LOCAL

### INSTALL

```
go get -u -v -ldflags="-s -w" github.com/tanakapayam/mtzdate
```

### ENVIRONMENT

```
export MTZDATE_TIMEZONES='San Francisco:America/Los_Angeles,UTC,München:Europe/Berlin,काठमाडौं:Asia/Kathmandu,東京:Asia/Tokyo'
export MTZDATE_FLAGS='San Francisco:US,München:DE,काठमाडौं:NP,東京:JP'
export MTZDATE_WORKDAYS='Mon,Tue,Wed,Thu,Fri'
export MTZDATE_GREEN_HOURS='8-17'
export MTZDATE_YELLOW_HOURS='7-8,17-18'
export MTZDATE_FAINT_HOURS='0-7,22-24'
export MTZDATE_FORMAT='dfc'
```

### RUN

One time:

```
mtzdate
```

Loop (^C to break):

```
mtzdate --loop
```

### HELP

```
Usage:
  mtzdate (-h | --version)
  mtzdate
  mtzdate --loop

Description
  This command-line utility displays Unix date in multiple time zones
  based on environment variables.

  With MTZDATE_LOOP=1 or --loop, mtzdate will refresh the screen once a second.
  Control-C will break the loop.

Options:
  -h, --help
  -l, --loop    # Loop until Control-C is trapped
  --version

Installation
  go get -u -v -ldflags="-s -w" github.com/tanakapayam/mtzdate

Environment
  Set MTZDATE_TIMEZONES to a comma-separated list of time zones. If desired, preface each time zone with a
  UTF-8-encoded city name or alias and a colon.

  To see emoji flags, set MTZDATE_FLAGS to a comma-separated map of UTF-8-encoded city names or aliases
  followd by two-letter country code -- separated by a colon.

  mtzdate defaults to coloring workhours green and to coloring pre- and post-workhours yellow. The behavior
  is controlled by the following environment variables (with their default values):

  export MTZDATE_WORKDAYS='Mon,Tue,Wed,Thu,Fri'
  export MTZDATE_GREEN_HOURS='8-17'
  export MTZDATE_YELLOW_HOURS='7-8,17-18'

  To opt out of the feature, set MTZDATE_WORKDAYS='':

  export MTZDATE_WORKDAYS=''

  MTZDATE_FORMAT can be set to a sequence of "d" (date), "f" (flag) and "c" (city) to signify the display
  format. (If unset, "dfc" is assumed.) Naturally, it's most meaningful if it's three letters, but there
  are no restrictions.

Examples
  $ export MTZDATE_TIMEZONES='America/Chicago,Europe/Paris'
  $ export MTZDATE_FLAGS='Chicago:US,Paris:FR'

  $ date
  Sun Jul 29 18:10:33 PDT 2018

  $ mtzdate
  Sun Jul 29 20:10:33 CDT   🇺🇸  Chicago
  Mon Jul 30 03:10:33 CEST  🇫🇷  Paris

  $ export MTZDATE_TIMEZONES='San Francisco:America/Los_Angeles,UTC,München:Europe/Berlin,काठमाडौं:Asia/Kathmandu,東京:Asia/Tokyo'
  $ export MTZDATE_FLAGS='San Francisco:US,München:DE,काठमाडौं:NP,東京:JP'

  $ mtzdate
  Sun Jul 29 18:10:33 PDT   🇺🇸  San Francisco
  Mon Jul 30 01:10:33 UTC   ☁️
  Mon Jul 30 03:10:33 CEST  🇩🇪  München
  Mon Jul 30 06:55:33 +0545 🇳🇵  काठमाडौं
  Mon Jul 30 10:10:33 JST   🇯🇵  東京

  $ export MTZDATE_FORMAT='fd'
  $ mtzdate
  🇺🇸  Sun Jul 29 18:10:33 PDT
  ☁️  Mon Jul 30 01:10:33 UTC
  🇩🇪  Mon Jul 30 03:10:33 CEST
  🇳🇵  Mon Jul 30 06:55:33 +0545
  🇯🇵  Mon Jul 30 10:10:33 JST

See Also
  /usr/share/zoneinfo
  https://github.com/tanakapayam/mtzdate
```

### TESTED ON

```
go version go1.10.3 darwin/amd64
```

## DOCKER

### BUILD

```
make docker-build
```

### PULL

```
docker pull tanakapayam/mtzdate
```

### RUN

One time:

```
docker run \
  --env MTZDATE_FLAGS='San Francisco:US,München:DE,काठमाडौं:NP,東京:JP' \
  --env MTZDATE_TIMEZONES='San Francisco:America/Los_Angeles,UTC,München:Europe/Berlin,काठमाडौं:Asia/Kathmandu,東京:Asia/Tokyo' \
  --tty \
  tanakapayam/mtzdate
```

Loop (^C to break):

```
docker run \
  --env MTZDATE_LOOP=1 \
  --env MTZDATE_FLAGS='San Francisco:US,München:DE,काठमाडौं:NP,東京:JP' \
  --env MTZDATE_TIMEZONES='San Francisco:America/Los_Angeles,UTC,München:Europe/Berlin,काठमाडौं:Asia/Kathmandu,東京:Asia/Tokyo' \
  --interactive \
  --tty \
  tanakapayam/mtzdate
```

## SEE ALSO

World Time Zones:

* [List of tz database time zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)

ISO Country Codes and Flag Emoji:

* [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
* [🎌 Flags](https://emojipedia.org/flags/)

## LICENSE

[MIT](https://github.com/tanakapayam/mtzdate/blob/master/LICENSE)
