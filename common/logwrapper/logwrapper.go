package logwrapper

import (
	"context"
	"dic_crawler/common/traceid"
	"flag"
	"fmt"
	"github.com/golang/glog"
)

// Init 初始化日志库
func Init(logDir string, logLevel int) {
	flag.Set("log_dir", logDir)
	flag.Set("v", fmt.Sprintf("%d", logLevel))
	flag.Set("logtostderr", "true")
	flag.Set("alsologtostderr", "true")
	flag.Parse()
}

// Flush 刷新日志
func Flush() {
	glog.Flush()
}

// Info 记录信息级别日志
func Info(ctx context.Context, args ...interface{}) {
	traceID := traceid.GetTraceID(ctx)
	glog.InfoDepth(1, append([]interface{}{fmt.Sprintf("[traceid: %s]", traceID)}, args...)...)
}

// Warning 记录警告级别日志
func Warning(ctx context.Context, args ...interface{}) {
	traceID := traceid.GetTraceID(ctx)
	glog.WarningDepth(1, append([]interface{}{fmt.Sprintf("[traceid: %s]", traceID)}, args...)...)
}

// Error 记录错误级别日志
func Error(ctx context.Context, args ...interface{}) {
	traceID := traceid.GetTraceID(ctx)
	glog.ErrorDepth(1, append([]interface{}{fmt.Sprintf("[traceid: %s]", traceID)}, args...)...)
}

// Fatal 记录致命错误级别日志并终止程序
func Fatal(ctx context.Context, args ...interface{}) {
	traceID := traceid.GetTraceID(ctx)
	glog.FatalDepth(1, append([]interface{}{fmt.Sprintf("[traceid: %s]", traceID)}, args...)...)
}

// V 返回一个是否启用的级别
func V(level int) glog.Verbose {
	return glog.V(glog.Level(level))
}
