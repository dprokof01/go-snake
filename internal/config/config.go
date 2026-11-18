package config

type Config struct {
	Size Size `yaml:"size"`
}

type Size struct {
	Col int `yaml:"col"`
	Row int `yaml:"row"`
}
