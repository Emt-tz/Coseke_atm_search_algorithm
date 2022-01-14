package main

import (
	"fmt"
	"sort"
)

type Atm struct {
	id        int
	latitude  float64
	longitude float64
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

/*
Fetch latitude and longitude using the gps
*/
func gps_location_fetch() (Location, error) {
	//scan gps
	scanned_gps_results := Location{-6.7, -3.2}
	return scanned_gps_results, nil
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
	all_atms_list["atm1"] = Atm{1, 32.9697, -96.80322}
	all_atms_list["atm2"] = Atm{2, 29.46786, -98.53506}
	all_atms_list["atm3"] = Atm{3, 32.9697, -95.80322}
	nearest_atms_list := make(map[int]Atm)

	if all_atms_list == nil {
		return nil, nil
	}
	//if not nil
	for _, atm := range all_atms_list {
		if currentUser.latitude < 2*(atm.latitude) {
			nearest_atms_list[atm.id] = atm
		} else {
			fmt.Println(distance(currentUser, atm, "K"))
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
		atm_distance := distance(currentUser, atm, "K")
		nearest_atms_list_distances[atmid] = UserToAtmDistance{atm: atm, distance: atm_distance}
	}
	return nearest_atms_list_distances, nil
}

/*
Finally connect the user to the nearest atm
*/
func Connect_user_to_nearest_atm(currentUser User) (UserToAtmDistance, error) {
	var closest_atm UserToAtmDistance
	nearest_atms_list_distances, err := Calculate_user_to_atms_distances(currentUser)
	if err != nil {
		return UserToAtmDistance{}, err
	}
	temp_slice := make(UserToAtmDistanceSlice, 0, len(nearest_atms_list_distances))
	for _, user_to_atm_pointer := range nearest_atms_list_distances {
		temp_slice = append(temp_slice, &user_to_atm_pointer)
	}
	sort.Sort(temp_slice) //sort our temp slice
	for _, user_atm_distance_value := range temp_slice {
		closest_atm = *user_atm_distance_value
	}
	return closest_atm, nil
}

func main() {
	//current user
	user := User{1, 32.9697, -96.80322}
	result, err := Connect_user_to_nearest_atm(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(
		"The closest atm found\nid: %v\nlatitude: %f\nlongitude: %f\ndistance: %f km\n", result.atm.id, result.atm.latitude, result.atm.longitude, result.distance)
}
