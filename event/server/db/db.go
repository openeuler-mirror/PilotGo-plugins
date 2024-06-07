package db

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"openeuler.org/PilotGo/PilotGo-plugin-event/config"
)

var InfluxDB *InfluxDBClient

type InfluxDBClient struct {
	Organization string
	Bucket       string
	DBClient     influxdb2.Client
}

func InfluxdbInit(conf *config.Influxd) {
	c := influxdb2.NewClient(conf.URL, conf.Token)

	InfluxDB = &InfluxDBClient{
		Organization: conf.Organization,
		Bucket:       conf.Bucket,
		DBClient:     c,
	}
}
