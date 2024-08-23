package main

import (
	"github.com/oschwald/geoip2-golang"
	"log"
	"targeting_service/internal/ads"
)

func main() {
	geoip, err := geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer geoip.Close()

	server := ads.NewServer(geoip)
	log.Fatalln(server.Listen())
}
