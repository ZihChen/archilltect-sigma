package structs

type Config struct {
	Name      string    `yaml:"name"`
	Port      int       `yaml:"port"`
	LogConfig LogConfig `yaml:"log"`
	GptConfig GptConfig `yaml:"gpt"`
}

type LogConfig struct {
	Level      string `yaml:"level"`
	FileName   string `yaml:"filename"`
	MaxSize    int    `yaml:"max_size"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"`
}

type GptConfig struct {
	BaseUrl string `yaml:"base_url"`
}
