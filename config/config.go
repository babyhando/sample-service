package config

type Config struct {
	Version string `mapstructure:"version"`
	Server  Server `mapstructure:"server"`
}

type Server struct {
	HttpPort int    `mapstructure:"http_port"`
	Host     string `mapstructure:"host"`
}
