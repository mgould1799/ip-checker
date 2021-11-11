package main

import "github.com/mgould1799/ip-checker/api"


func main() {
	server := api.NewServer("./ipChecker/GeoLite2-Country_20211109/GeoLite2-Country.mmdb", 8080)
	server.Start() 
}