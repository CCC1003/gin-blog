package config

type Logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     string `yaml:"show-line"`
	LogInConsole string `yaml:"log-in-console"`
}
