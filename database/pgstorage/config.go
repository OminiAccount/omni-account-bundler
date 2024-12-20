package pgstorage

// Config struct
type Config struct {
	// Database name
	Name string `mapstructure:"name"`

	// User name
	User string `mapstructure:"user"`

	// Password of the user
	Password string `mapstructure:"password"`

	// Host address
	Host string `mapstructure:"host"`

	// Port Number
	Port string `mapstructure:"port"`

	// EnableLog
	EnableLog bool `mapstructure:"enable-log"`

	LogFile string `mapstructure:"log-file"`

	LogLevel string `mapstructure:"log-level" jsonschema:"enum=debug,enum=info,enum=warn,enum=error"`

	// Unit in milliseconds
	LogSlowTime int64 `mapstructure:"log-slow-time" jsonschema:"enum=debug,enum=info,enum=warn,enum=error"`

	// MaxConns is the maximum number of connections in the pool.
	MaxConns int `mapstructure:"max-conns"`
}
