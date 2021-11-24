package domain

type Configuration struct {
	InfluxDb *InfluxDb `yaml:"influxdb"`
}

type InfluxDb struct {
	Url string `yaml:"Url"`
	Port int `yaml:"Port"`
	Measurement string `yaml:"Measurement"`
	Database string `yaml:"Database"`
	RetentionPolicy string `yaml:"RetentionPolicy"`
}
