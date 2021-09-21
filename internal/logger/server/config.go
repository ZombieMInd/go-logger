package server

type Config struct {
	Name     string `envconfig:"NAME" default:"logger"`
	BindAddr string `envconfig:"BIND_ADDR" default:"127.0.0.1:8080"`
}