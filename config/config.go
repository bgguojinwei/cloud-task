package config

type Server struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
}

type System struct {
	Version string `mapstructure:"version" json:"version" yaml:"version"`
}
