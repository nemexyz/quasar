package main

import (
	"fmt"
	"os"
	"quasar/usecase"
	"quasar/util"
	"strconv"
)

func main() {
	args := os.Args[1:]
	command := args[0]

	if command == "location" {
		distKen, _ := strconv.ParseFloat(args[1], 64)
		distSky, _ := strconv.ParseFloat(args[2], 64)
		distSato, _ := strconv.ParseFloat(args[3], 64)
		location := usecase.GetLocation(distKen, distSky, distSato)
		fmt.Printf("Enemy location: (%f, %f)\n", location[0], location[1])
	} else if command == "message" {
		data := util.GetDataToArraysString(args[1])
		fmt.Printf("Enemy message decoded: '%s'\n", usecase.GetMessage(data))
	} else if command == "server" {
		Server()
	} else {
		fmt.Printf("Command: %s not available, use 'message' 'location' or 'server'\n", command)
	}
}
