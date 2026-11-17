package config

type Config struct {
	Size Size `yaml:"size"`
}

type Size struct {
	Row int `yaml:"row"`
	Col int `yaml:"col"`
}
