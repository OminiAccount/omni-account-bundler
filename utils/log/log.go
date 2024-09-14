package log

import (
	"github.com/ethereum/go-ethereum/log"
	"golang.org/x/exp/slog"
	"os"
)

// SetGlobalLogHandler sets the log handles as the handler of the global default logger.
// The usage of this logger is strongly discouraged,
// as it does makes it difficult to distinguish different services in the same process, e.g. during tests.
// Geth and other components may use the global logger however,
// and it is thus recommended to set the global log handler to catch these logs.
func SetGlobalLogHandler(h slog.Handler) {
	log.SetDefault(log.NewLogger(h))
}

func SetupDefaults() {
	SetGlobalLogHandler(log.NewTerminalHandlerWithLevel(os.Stdout, log.LevelDebug, true))
}
