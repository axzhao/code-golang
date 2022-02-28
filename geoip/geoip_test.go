package geoip

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
	"net/http"
	"strings"
)

func getIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func ExampleIP() {
	db, err := geoip2.Open("/tmp/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP("34.92.133.165")
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("city name: %v\n", record.City.Names["en"])
	fmt.Printf("city name: %v\n", record.City.Names["zh-CN"])
	if len(record.Subdivisions) > 0 {
		fmt.Printf("subdivision name: %v\n", record.Subdivisions[0].Names["en"])
		fmt.Printf("subdivision name: %v\n", record.Subdivisions[0].Names["zh-CN"])
	}
	fmt.Printf("country name: %v\n", record.Country.Names["en"])
	fmt.Printf("country name: %v\n", record.Country.Names["zh-CN"])
	fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
	fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
	fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)
	// Output:
}
