package pgstorage

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
)

type logger struct {
	log      *Logger
	slowTime int64
}

func (l logger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	queryTimeStr := fmt.Sprintf("%s", data["time"])
	queryTime, _ := time.ParseDuration(queryTimeStr)
	if queryTime.Milliseconds() < l.slowTime {
		return
	}
	m := fmt.Sprintf("%s %v", msg, data)

	switch level {
	case pgx.LogLevelInfo:
		l.log.Info(m)
	case pgx.LogLevelWarn:
		l.log.Warn(m)
	case pgx.LogLevelError:
		l.log.Error(m)
	default:
		m = fmt.Sprintf("%s %s %v", level.String(), msg, data)
		l.log.Debug(m)
	}
}
