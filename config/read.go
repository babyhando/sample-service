package config

import (
	"path/filepath"

	"github.com/spf13/viper"
)

func ReadGeneric[T any](cfgPath string) (T, error) {
	var cfg T
	fullAbsPath, err := absPath(cfgPath)
	if err != nil {
		return cfg, err
	}

	// configPath := filepath.Dir(fullAbsPath)
	// viper.AddConfigPath(configPath)
	// configType := strings.TrimPrefix(filepath.Ext(fullAbsPath), ".")
	// viper.SetConfigType(configType)
	// configFile := strings.TrimSuffix(filepath.Base(fullAbsPath), filepath.Ext(fullAbsPath))
	// viper.SetConfigName(configFile)

	viper.SetConfigFile(fullAbsPath)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return cfg, err
	}

	return cfg, viper.Unmarshal(&cfg)
}

func ReadStandard(cfgPath string) (Config, error) {
	return ReadGeneric[Config](cfgPath)
}

func absPath(cfgPath string) (string, error) {
	if !filepath.IsAbs(cfgPath) {
		return filepath.Abs(cfgPath)
	}
	return cfgPath, nil
}

func MustReadStandard(configPath string) Config {
	cfg, err := ReadStandard(configPath)
	if err != nil {
		panic(err)
	}
	return cfg
}
