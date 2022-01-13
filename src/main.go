package main

import (
	"fmt"
	"math"
	"sort"
)

type Atm struct {
	id         int
	latitude   float64
	longitude  float64
}

type User struct {
	id        int
	latitude  float64
	longitude float64
}

type Location struct {
	latitude  float64
	longitude float64
}

type UserToAtmDistanceSlice []*UserToAtmDistance

type UserToAtmDistance struct {
	atm      Atm
	distance float64
}

var defaultMeasurementDistance float64 = 200

func main() {
	//current user
	user := User{1, 32.9697, -96.80322}
	fmt.Println(Connect_user_to_nearest_atm(user))
}

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
End External sorting functions
*/

/*
Fetch latitude and longitude using the gps
*/
func gps_location_fetch() (Location, error) {
	//scan gps
	scanned_gps_results := Location{-6.7, -3.2}
	return scanned_gps_results, nil
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

/*
Fetch the current user location and return a User type struct
*/
func Get_current_user_location(userid int) (User, error) {
	//user lat and long
	userlocation, err := gps_location_fetch()
	if err != nil {
		return User{}, err
	}
	return User{userid, userlocation.latitude, userlocation.longitude}, nil
}

/*
This helps to filter and prevent search for entire database of atms thus list only atms nearer the location
*/
func Get_all_atms_based_on_current_location(currentUser User) (map[int]Atm, error) {
	all_atms_list := make(map[string]Atm)

  //generate sample data for atm list
  all_atms_list["atm1"] = Atm{1,32.9697, -96.80322}
  all_atms_list["atm2"] = Atm{2,29.46786, -98.53506}
  all_atms_list["atm3"] = Atm{3,32.9697, -95.80322}
	nearest_atms_list := make(map[int]Atm)

	if all_atms_list == nil {
		return nil, nil
	}
	//if not nil
	for _, atm := range all_atms_list {
		if currentUser.latitude > 2*(atm.latitude) {
			nearest_atms_list[atm.id] = atm
		}else{
      fmt.Println(distance(currentUser, atm))
    }
	}
	return nearest_atms_list, nil
}

/*
Deeply evaluate the distance between each atm and the user current location so as to provide the best and closest distance atm
*/
func Calculate_user_to_atms_distances(currentUser User) (map[int]UserToAtmDistance, error) {
	nearest_atms_list, err := Get_all_atms_based_on_current_location(currentUser)
	nearest_atms_list_distances := make(map[int]UserToAtmDistance) //for each atm map atm id with distance
	if err != nil {
		return nil, err
	}
	for atmid, atm := range nearest_atms_list {
		atm_distance := distance(currentUser, atm)
		nearest_atms_list_distances[atmid] = UserToAtmDistance{atm: atm, distance: atm_distance}
	}
	return nearest_atms_list_distances, nil
}

/*
Finally connect the user to the nearest atm
*/
func Connect_user_to_nearest_atm(currentUser User) (Atm, error) {
	var closest_atm Atm
	nearest_atms_list_distances, err := Calculate_user_to_atms_distances(currentUser)
	if err != nil {
		return Atm{}, err
	}

	s := make(UserToAtmDistanceSlice, 0, len(nearest_atms_list_distances))

	for _, d := range s {
		s = append(s, d)
	}

	sort.Sort(s)
	for _, d := range s {
		fmt.Printf("%+v\n", *d)
	}
	return closest_atm, nil
}
