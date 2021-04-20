# GIS Utils
This package is a go package that is utilities that are used by wwum 2020 CLI for the GIS functions that might
want to be used in other programs as well. We decided this could then be a 
standalone package so that it can be added to many others with ease.

## Functions
There are several functions in this package that help with spatial operations

### Haversin Distance
Haversin Distance Formula for great arc distance on a sphere with accuracy for small distances. The formula
returns the meters between two points using decimal latitude and longitude

### Inverse Distance Weighted
Produces the inverse distance weighted results of a slice passed to it and returns
a normalized 4 digit percentage of those values.