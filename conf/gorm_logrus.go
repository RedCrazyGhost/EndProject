/**
 @author: RedCrazyGhost
 @date: 2023/4/6

**/

package conf

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

var lo logrus.Logger

type GromAndLogrusLog struct {
	logger.Config
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

func New() *GromAndLogrusLog {

	log := &GromAndLogrusLog{}
	log.Config.LogLevel = logger.Info
	return log
}

func (l *GromAndLogrusLog) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info print info
func (l *GromAndLogrusLog) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		Log.WithField("log.file", utils.FileWithLineNum()).Printf(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Warn print warn messages
func (l *GromAndLogrusLog) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		Log.WithField("log.file", utils.FileWithLineNum()).Printf(l.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Error print error messages
func (l *GromAndLogrusLog) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		Log.WithField("log.file", utils.FileWithLineNum()).Printf(l.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

// Trace print sql message
func (l *GromAndLogrusLog) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			Log.WithField("log.file", utils.FileWithLineNum()).Printf(l.traceErrStr, err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			Log.WithField("log.file", utils.FileWithLineNum()).Printf(l.traceErrStr, err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			Log.WithField("log.file", utils.FileWithLineNum()).Printf(l.traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			Log.WithField("log.file", utils.FileWithLineNum()).Printf(l.traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.LogLevel == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			Log.WithField("log.file", utils.FileWithLineNum()).Printf(l.traceStr, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			Log.WithField("log.file", utils.FileWithLineNum()).Printf(l.traceStr, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
