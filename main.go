package main

import "github.com/SamPariatIL/weather-wrapper/cmd"

// @title Weather Wrapper API
// @version 1.0
// @description This is a wrapper for the OpenWeatherMap API.
// @host localhost:8181
// @BasePath /api/v1
func main() {
	cmd.RunServer()
}
