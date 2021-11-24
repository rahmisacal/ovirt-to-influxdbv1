package infrastructure

import (
	"../domain"
	"encoding/json"
	"fmt"
)

func MapMetric(url string, token string) {

	var (
		data       = domain.Json{}
		statistics = domain.Statistics{}
	)
	_ = json.Unmarshal(GetMetric(url, token), &data) // Get Hosts and map json

	var ids [50]string //TODO

	for i, metrics := range data.Host {

		ids[i] = metrics.Id //Append Host Ids to array
		for _, id := range ids {
			host := url + "/" + id + "/" + "statistics"
			_ = json.Unmarshal(GetMetric(host, token), &statistics) //Get Hosts Statistics and map statistics

			for _, statistic := range statistics.Statistic {
				if statistic.Name == "memory.used" || statistic.Name == "cpu.current.system" {
					value := fmt.Sprintf("%.2f", statistic.Values.Value[0].Datum)
					Influxdb(
						metrics.Name,
						statistic.Name,
						value,
					)
				}
			}
		}
	}
}
