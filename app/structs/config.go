package structs

type Config struct {
	Name      string     `mapstructure:"name"`
	Port      int        `mapstructure:"port"`
	LogConfig *LogConfig `mapstructure:"log"`
	GptConfig *GptConfig `mapstructure:"gpt"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type GptConfig struct {
	BaseUrl string `mapstructure:"base_url"`
}
