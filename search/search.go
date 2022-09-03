package search

import "math"

type Request struct {
	Username string `json:"username"`
	Path     Path   `json:"path"`
}

type Path struct {
	StartLocation Location `json:"start_location"`
	EndLocation   Location `json:"end_location"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64) float64 {
	radlat1 := float64(math.Pi * lat1 / 180)
	radlat2 := float64(math.Pi * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515 * 1.609344
	return dist
}

func distanceBetweenTwoLocations(loc1 Location, loc2 Location) float64 {
	return distance(loc1.Latitude, loc1.Longitude, loc2.Latitude, loc2.Longitude)
}

func (r Request) IsValidPartner(other Request) bool {
	return distanceBetweenTwoLocations(r.Path.StartLocation, other.Path.StartLocation) < 1 && distanceBetweenTwoLocations(r.Path.EndLocation, other.Path.EndLocation) < 1
}

func GetClosestLocations(plan1 Request, plan2 Request) Path {
	startLocation := plan1.Path.StartLocation
	endLocation := plan1.Path.EndLocation
	minDist := distanceBetweenTwoLocations(plan1.Path.StartLocation, plan1.Path.EndLocation)
	if distanceBetweenTwoLocations(plan1.Path.StartLocation, plan2.Path.EndLocation) < minDist {
		startLocation = plan1.Path.StartLocation
		endLocation = plan2.Path.EndLocation
		minDist = distanceBetweenTwoLocations(plan1.Path.StartLocation, plan2.Path.EndLocation)
	}
	if distanceBetweenTwoLocations(plan2.Path.StartLocation, plan1.Path.EndLocation) < minDist {
		startLocation = plan2.Path.StartLocation
		endLocation = plan1.Path.EndLocation
		minDist = distanceBetweenTwoLocations(plan2.Path.StartLocation, plan1.Path.EndLocation)
	}
	if distanceBetweenTwoLocations(plan2.Path.StartLocation, plan2.Path.EndLocation) < minDist {
		startLocation = plan2.Path.StartLocation
		endLocation = plan2.Path.EndLocation
	}
	return Path{StartLocation: startLocation, EndLocation: endLocation}
}
