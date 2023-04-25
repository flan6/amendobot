package config

import "github.com/spf13/viper"

func init() {
	viper.SetConfigFile("config/.env")
	viper.ReadInConfig()
}

func DiscordToken() string {
	return viper.GetString("DISCORD_TOKEN")
}
