package main

import (
	geoIp2 "github.com/oschwald/geoip2-golang"
	"log"
	"net"
	"os"
)

type IpCountry struct {
	IsoCode string            `json:"isoCode"`
	Names   map[string]string `json:"names"`
}
type IpCity struct {
	GeoNameID uint              `json:"geoNameId"`
	Names     map[string]string `json:"names"`
}
type IpInfo struct {
	Ip        string    `json:"ip"`
	Country   IpCountry `json:"country"`
	City      IpCity    `json:"city"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	TimeZone  string    `json:"timeZone"`
}

var (
	geoClient *GeoClient
)

func InitGeoIp2Client() {
	geoFilePath := os.Getenv("GEO_FILE")
	geoClient = &GeoClient{}
	var err error
	geoClient.IpInfoReader, err = geoIp2.Open(geoFilePath)
	if err != nil {
		log.Fatal("初始化IP数据库错误")
	}
}

func (client *GeoClient) GeoIp2GetIpInfo(ipStr string) (ipInfo *IpInfo, err error) {
	ip := net.ParseIP(ipStr)
	var record *geoIp2.City
	record, err = client.IpInfoReader.City(ip)
	if err != nil {
		return
	}
	ipInfo = &IpInfo{
		Ip: ipStr,
	}
	ipInfo.Ip = ipStr
	ipInfo.Country = IpCountry{IsoCode: record.Country.IsoCode, Names: record.Country.Names}
	ipInfo.City = IpCity{GeoNameID: record.City.GeoNameID, Names: record.City.Names}
	ipInfo.Latitude = record.Location.Latitude
	ipInfo.Longitude = record.Location.Longitude
	ipInfo.TimeZone = record.Location.TimeZone
	return
}
