# ovirt-influxdb exporter
A Simple Application Using Golang

## About : ovirt-influxdb exporter
The App is used golang http client, authenticated to ovirt-rest-api with token, received metrics of the hosts finally this application is sent this metrics to influxdb every 3 minutes

- golang 1.17.3
- influxdb-client-v1(influxdb v1.7.9)

## Usage

 - Set environment deployment.yaml (url)
 - sure that access to influxdb:port and target Ovirt-rest-api from sourceIp