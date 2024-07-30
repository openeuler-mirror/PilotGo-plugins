package db

import (
	"context"
	"fmt"

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

func Query(measurement, start, end string) error {

	query := fmt.Sprintf("from(bucket:%s)|> range(start: -1h) |> filter(fn: (r) => r._measurement == %s)", InfluxDB.Bucket, measurement)

	queryAPI := InfluxDB.DBClient.QueryAPI(InfluxDB.Organization)

	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return err
	}

	for result.Next() {
		if result.TableChanged() {
			fmt.Printf("table: %s\n", result.TableMetadata().String())
		}

		fmt.Printf("time: %v, field: %v, value: %v\n", result.Record().Time().Format("2006-01-02 15:04:05"), result.Record().Field(), result.Record().Value())

	}

	if result.Err() != nil {
		fmt.Printf("query parsing error: %s\n", result.Err().Error())
	}

	return nil
}
