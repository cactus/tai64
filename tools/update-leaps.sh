#!/usr/bin/env bash

# sometimes sites go offline, so list a few options.
# just use one.
{
    #curl -s 'http://maia.usno.navy.mil/ser7/tai-utc.dat'
    #curl -s 'https://cdf.gsfc.nasa.gov/html/CDFLeapSeconds.txt'
    #curl -s 'https://hpiers.obspm.fr/iers/bul/bulc/Leap_Second.dat'| grep -v '^#'| awk '{print $4" "$3" "$2" "$5}'
} > leaps.dat

go run ./tools/generate.go -input leaps.dat -output offsets.go -pkg tai64
go fmt offsets.go
