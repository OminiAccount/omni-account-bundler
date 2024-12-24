package state

import "github.com/OAB/utils/apitypes"

type Config struct {
	HisInterval apitypes.Duration `mapstructure:"his-interval"`
}
