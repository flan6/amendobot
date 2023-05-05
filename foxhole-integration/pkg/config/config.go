package config

import "github.com/spf13/viper"

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.SetDefault("SHARD1_URL", "https://war-service-live.foxholeservices.com/api/")
	viper.SetDefault("BADGER_PATH", "/tmp/badger")
}

func WarApiURL() string {
	return viper.GetString("SHARD1_URL")
}

func BadgerPath() string {
	return viper.GetString("BADGER_PATH")
}
