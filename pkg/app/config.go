package app

import "flag"

type Config struct {
	Host        string
	Port        string
	DbHost      string
	DiscountUrl string
}

func ParseConfig() *Config {
	conf := &Config{}

	flag.StringVar(&conf.Host, "host", "localhost", "Server host")
	flag.StringVar(&conf.Port, "port", "8081", "Server port")
	flag.StringVar(&conf.DbHost, "dbHost", "file:/tmp/catalog.db?_mutex=full&_cslike=false", "Server dbHost")
	flag.StringVar(&conf.DiscountUrl, "discountUrl", "http://localhost:8081/static/discount.csv", "Url for discount.csv")
	flag.Parse()

	return conf
}
