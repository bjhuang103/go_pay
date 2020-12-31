package logs

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	. "gorm.io/gorm/logger"
	"time"
)

// 实现gorm的log
type GLogImpl struct {
}

var defaultLog = GLogImpl{}

func (l *GLogImpl) LogMode(LogLevel) Interface {
	return l
}

func (l *GLogImpl) Info(ctx context.Context, msg string, data ...interface{}) {
	logs := logx.WithContext(ctx)
	logs.Infof(msg, data)
}

func (l *GLogImpl) Warn(ctx context.Context, msg string, data ...interface{}) {
	logs := logx.WithContext(ctx)
	logs.Infof(msg, data)
}

func (l *GLogImpl) Error(ctx context.Context, msg string, data ...interface{}) {
	logs := logx.WithContext(ctx)
	logs.Errorf(msg, data)
}

func (l *GLogImpl) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
}

func CtxInfo(ctx context.Context, msg string, data ...interface{}) {
	defaultLog.Info(ctx, msg, data)
}

func CtxWarn(ctx context.Context, msg string, data ...interface{}) {
	defaultLog.Warn(ctx, msg, data)
}

func CtxError(ctx context.Context, msg string, data ...interface{}) {
	defaultLog.Error(ctx, msg, data)
}
