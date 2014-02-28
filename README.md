## Geo
A library in Go dealing with earth coordinate system (geodetic) calculations.

[![Build Status](https://travis-ci.org/mattfarina/geo-go.png?branch=master)](https://travis-ci.org/mattfarina/geo-go) [![GoDoc](https://godoc.org/github.com/mattfarina/geo-go?status.png)](https://godoc.org/github.com/mattfarina/geo-go) [![Coverage Status](https://coveralls.io/repos/mattfarina/geo-go/badge.png?branch=master)](https://coveralls.io/r/mattfarina/geo-go?branch=master)

## Locations
Locations are individual locations on the Earth.

    location := NewLocation(lat, long)

From locations you can do things like calculate the distance between two of them. For example,

    l1 := NewLocation(lat1, long1)
    l2 := NewLocation(lat2, long2)
    distance := l1.Distance(l2)

You can also us a location to find a range or longitudes and latitudes. This is useful for finding the corners on the little rectangle screens we look at too often. For example,

    location := NewLocation(42.7186, -84.468466)
    latMin, latMax := location.LatitudeRange(10000)
    longMin, longMax := location.LongitudeRange(20000)

There are other operations you can find in the `location.go` code.

## Earth
These locations are on the Earth and information for the earth is stored on Earth. For example, the method `EarthRadius` will return the radius of the Earth at a specified latitude. This takes into account that the earth is an irregular ellipsoid.

There are other operations that can be performed on the Earth. For more details see `earth.go`.
