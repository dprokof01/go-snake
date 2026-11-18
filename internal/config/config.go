package config

type Config struct {
	Size Size `yaml:"size"`
}

type Size struct {
	Col int `yaml:"col" validate:"min=3"`
	Row int `yaml:"row" validate:"min=3"`
}
