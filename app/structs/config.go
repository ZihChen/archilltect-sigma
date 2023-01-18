package structs

type Config struct {
	Name       string      `mapstructure:"name"`
	Port       int         `mapstructure:"port"`
	LineConfig *LineConfig `mapstructure:"line"`
	GptConfig  *GptConfig  `mapstructure:"gpt"`
}

type GptConfig struct {
	BaseUrl string `mapstructure:"base_url"`
	Key     string `mapstructure:"key"`
}

type LineConfig struct {
	ChannelToken  string `mapstructure:"channel_token"`
	ChannelSecret string `mapstructure:"channel_secret"`
}
