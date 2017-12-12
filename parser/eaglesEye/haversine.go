package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"math"
	"os"
)

const kmtomiles = float64(0.621371192)
const earthRadius = float64(6371)

//{"type":"node","id":124682,"lat":45.2466902,"lon":5.6555821,"tags":{}}
type Node struct {
	Type []string `json:"type"`
	ID   int      `json:"id"`
	Lat  float64  `json:"lat"`
	Lon  float64  `json:"lon"`
}

func main() {
	lf := os.Args[1]
	locationFromJson := getNode(lf)
	res := Node{}
	json.Unmarshal([]byte(locationFromJson), &res)
	locationFromID := res.ID
	locationFromLat := res.Lat
	locationFromLon := res.Lon

	lt := os.Args[2]
	locationToJson := getNode(lt)
	res2 := Node{}
	json.Unmarshal([]byte(locationToJson), &res2)
	locationToID := res2.ID
	locationToLat := res2.Lat
	locationToLon := res2.Lon

	// Use haversine to get the resulting diatance between the two values
	var distance = Haversine(locationFromLat, locationFromLon, locationToLat, locationToLon)
	// We wish to use miles so will alter the resulting distance
	var distancemiles = distance * kmtomiles

	fmt.Printf("The distance between %s and %s is %.02f miles as the crow flies", locationFromID, locationToID, distancemiles)
}

/*
 * The haversine formula will calculate the spherical distance as the crow flies
 * between lat and lon for two given points in km
 */
func Haversine(lonFrom float64, latFrom float64, lonTo float64, latTo float64) (distance float64) {
	var deltaLat = (latTo - latFrom) * (math.Pi / 180)
	var deltaLon = (lonTo - lonFrom) * (math.Pi / 180)

	var a = math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(latFrom*(math.Pi/180))*math.Cos(latTo*(math.Pi/180))*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance = earthRadius * c

	return
}

func getNode(id string) string {

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0, // use default DB
	})

	val2, err := client.Get(id).Result()
	if err == redis.Nil {
		fmt.Println(id, " does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println(id, val2)
	}
	return id
}
