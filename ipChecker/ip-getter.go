package ipChecker

import (
	"net"

	log "github.com/sirupsen/logrus"

	"github.com/oschwald/geoip2-golang"
)

type IpChecker struct {
	db *geoip2.Reader
}

func NewIpChecker(dbFile string) (*IpChecker) {
	ipChecker := IpChecker{}
	var err error
	ipChecker.db, err = geoip2.Open(dbFile)
	if err != nil {
		log.Panic(err)
	}

	return &ipChecker
}

func (ipC IpChecker) StopIpChecker() {
	ipC.db.Close()
}

func(ipC IpChecker) GetCountryFromIp(ip string) (*geoip2.Country, error){
	parsedIp := net.ParseIP(ip)
	countryRecord, err := ipC.db.Country(parsedIp)
	if err != nil {
		return nil, err 
	}
	return countryRecord, nil	
}

func(ipC IpChecker) InWhiteList(whiteListOfCountries[]string, ip string) (bool, error){
	ipsCountryInfo, err := ipC.GetCountryFromIp(ip)
	if err != nil {
		return false, err 
	}

	return stringInSlice(ipsCountryInfo.RegisteredCountry.Names["en"], whiteListOfCountries), nil 
}

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}
