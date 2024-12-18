package config

import (
	"bytes"
	"github.com/OAB/etherman"
	"github.com/OAB/jsonrpc"
	"github.com/OAB/pool"
	"github.com/OAB/state"
	"github.com/OAB/utils/log"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper"
	"github.com/urfave/cli/v2"
	"path/filepath"
	"strings"
)

const FlagCfg = "config"

// Config represents the `env.toml` file used to configure the processor
type Config struct {
	Ethereum etherman.Config    `mapstructure:"ethereum"`
	JsonRpc  jsonrpc.RpcsConfig `mapstructure:"jsonrpc"`
	Pool     pool.Config        `mapstructure:"pool"`
	State    state.Config       `mapstructure:"state"`
	Log      log.Config         `mapstructure:"log"`
}

// Default parses the default configuration values.
func Default() (*Config, error) {
	var cfg Config
	viper.SetConfigType("toml")

	err := viper.ReadConfig(bytes.NewBuffer([]byte(DefaultValues)))
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&cfg, viper.DecodeHook(mapstructure.TextUnmarshallerHookFunc()))
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

// Load loads the configuration
func Load(ctx *cli.Context) (*Config, error) {
	cfg, err := Default()
	if err != nil {
		return nil, err
	}
	configFilePath := ctx.String(FlagCfg)
	if configFilePath != "" {
		dirName, fileName := filepath.Split(configFilePath)

		fileExtension := strings.TrimPrefix(filepath.Ext(fileName), ".")
		fileNameWithoutExtension := strings.TrimSuffix(fileName, "."+fileExtension)

		viper.AddConfigPath(dirName)
		viper.SetConfigName(fileNameWithoutExtension)
		viper.SetConfigType(fileExtension)
	}
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("OMNI_BUNDLER")
	err = viper.ReadInConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if ok {
			log.Infof("config file not found")
		} else {
			log.Infof("error reading config file: ", err)
			return nil, err
		}
	}

	decodeHooks := []viper.DecoderConfigOption{
		// this allows arrays to be decoded from env var separated by ",", example: MY_VAR="value1,value2,value3"
		viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(mapstructure.TextUnmarshallerHookFunc(), mapstructure.StringToSliceHookFunc(","))),
	}

	err = viper.Unmarshal(&cfg, decodeHooks...)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
