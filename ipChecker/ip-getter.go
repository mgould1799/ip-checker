package ipChecker

import (
	"log"

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

