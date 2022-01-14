# Coseke_atm_search_algorithm
 A simple user nearest atm search algorithm in Golang

# language implementation

1. Golang (compiled language)

# Necessary Files
1. helpers.go
    a. Contains a function `distance` necessary to calculate distance between two latitudes and longitudes 
    b. Containes sorting functions based on the `golang` sorting library

2. main.go
    a. Contains detailed algorithm function and the main function to execute the code

# Algorithm design 

# Analysis:

1. Each Atm is represented as struct with type of `id, latitude and longitude`
2. A user is represented as a struct with type of `id, latitude and longitude` 
3. A user at a certain `location (l)` is far away distance (x) from the nearest atm
4. A user scans the current `location(l)` and gets results of the closest atm which is a short `distance (x)` away from the current user `location (l)`

# Higher level Algorithm:

1. Get the user location
2. Find the shortest distance atm from the user
3. Give the user results about the closest atm

# Detailed Algorithm:

1. Get the user location
    a. Scan the users gps to get the current latitude and longitude
    b. if not present abort with an error and exit
    c. if present scan the database with atms whose latitude and longitude are 2 times less the latitude and longitude of the current user (to filter out some very far away atms)

2. Finding the shortest distance atm
    a. Sort the filter atm list starting with the one with the smallest calculated distance value in kilometers
    b. Return the atm with min value of distance from the current user

3. Results 
    a. Return the user with the details of the closest atm


# Usage

1. Install the Golang sdk or use the go playground online
2. Clone the repository 
3. Through the command line `cd src` into src directory of the folder "Coseke_atm_search_algorithm"
4. Then execute the command `go run main.go` to see test results
