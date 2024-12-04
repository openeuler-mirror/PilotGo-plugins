/*
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugins licensed under the Mulan Permissive Software License, Version 2.
 * See LICENSE file for more details.
 * Author: zhanghan2021 <zhanghan@kylinos.cn>
 * Date: Fri Jun 7 17:32:09 2024 +0800
 */
package db

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"time"

	"gitee.com/openeuler/PilotGo-plugins/event/sdk"
	"gitee.com/openeuler/PilotGo/sdk/logger"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"openeuler.org/PilotGo/PilotGo-plugin-event/config"
)

var InfluxDB *InfluxDBClient

type InfluxDBClient struct {
	Organization string
	Bucket       string
	Measurement  string
	DBClient     influxdb2.Client
}

func InfluxdbInit(conf *config.Influxd) {
	c := influxdb2.NewClient(conf.URL, conf.Token)

	InfluxDB = &InfluxDBClient{
		Organization: conf.Organization,
		Bucket:       conf.Bucket,
		Measurement:  conf.Measurement,
		DBClient:     c,
	}
}

func Query(start, stop string, filterTagKey string) (interface{}, error) {
	query := fmt.Sprintf(`
	from(bucket:"%s")
	|> range(start: %s, stop: %s)
	|> filter(fn: (r) => r._measurement == "%s")
	`, InfluxDB.Bucket, start, stop, InfluxDB.Measurement)
	if filterTagKey != "" {
		query += fmt.Sprintf(`|> filter(fn: (r) => r.msg_type == "%s")`, filterTagKey)
	}

	queryAPI := InfluxDB.DBClient.QueryAPI(InfluxDB.Organization)
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return result, err
	}
	defer result.Close()

	var queryResults []map[string]interface{}
	for result.Next() {
		tags := make(map[string]string)
		for k, v := range result.Record().Values() {
			if k != "_measurement" && k != "_field" && k != "_value" && k != "_time" && k != "result" {
				if tagValue, ok := v.(string); ok {
					tags[k] = tagValue
				}
			}
		}
		queryResults = append(queryResults, map[string]interface{}{
			"value":    processValue(result.Record().Value()),
			"msg_type": tags["msg_type"],
			"time":     tags["timestamp"],
		})
	}

	sort.Slice(queryResults, func(i, j int) bool {
		timeFormat := "2006-01-02 15:04:05.999999999 -0700 MST"
		time1, err1 := time.Parse(timeFormat, queryResults[i]["time"].(string))
		time2, err2 := time.Parse(timeFormat, queryResults[j]["time"].(string))

		if err1 != nil || err2 != nil {

			return i < j
		}
		return time1.After(time2)
	})
	if result.Err() != nil {
		return nil, fmt.Errorf("查询数据出错: %v", result.Err())
	}
	return queryResults, nil
}

func WriteToDB(MessageData string) error {
	writeAPI := InfluxDB.DBClient.WriteAPIBlocking(InfluxDB.Organization, InfluxDB.Bucket)

	var msg sdk.MessageData
	err := json.Unmarshal([]byte(MessageData), &msg)
	if err != nil {
		logger.Error("解析数据出错: %v", err.Error())
		return err
	}

	tags := map[string]string{
		"msg_type":  msg.MessageType,
		"timestamp": msg.TimeStamp.String(),
	}
	fields := map[string]interface{}{
		"metadata": msg.Data,
	}
	point := write.NewPoint(InfluxDB.Measurement, tags, fields, msg.TimeStamp)

	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
		logger.Error("写入数据出错: %v", err.Error())
		return err
	}
	return nil
}

func processValue(v interface{}) interface{} {
	value := v.(string)
	re := regexp.MustCompile(`(\w+):([^\s\]]+)`)
	matches := re.FindAllStringSubmatch(value, -1)

	result := make(map[string]string)
	for _, match := range matches {
		result[match[1]] = match[2]
	}
	if len(matches) > 0 {
		result := make(map[string]string)
		for _, match := range matches {
			result[match[1]] = match[2]
		}
		// return result
		jsonBytes, _ := json.Marshal(result)
		return string(jsonBytes)
	}
	return value
}
