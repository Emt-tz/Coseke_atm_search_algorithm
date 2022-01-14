package main

import "math"

/*
External sorting functions using built in sort function
*/
// Len is part of sort.Interface.
func (d UserToAtmDistanceSlice) Len() int {
	return len(d)
}

// Swap is part of sort.Interface.
func (d UserToAtmDistanceSlice) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// Less is part of sort.Interface. We use count as the value to sort by
func (d UserToAtmDistanceSlice) Less(i, j int) bool {
	return d[i].distance < d[j].distance
}

/*
calculate distance between two latitudes and longitudes
return by default distance in miles which is unit
*/
func distance(origin User, destination Atm, unit ...string) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * origin.latitude / 180)
	radlat2 := float64(PI * destination.latitude / 180)

	theta := float64(origin.longitude - destination.latitude)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}
