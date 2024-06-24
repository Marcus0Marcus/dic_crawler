package traceid

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type key int

const traceIDKey key = iota

// NewTraceID 生成新的 traceid
func NewTraceID() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%016x", rand.Int63())
}

// WithTraceID 将 traceid 设置到 context 中
func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}

// GetTraceID 从 context 中获取 traceid
func GetTraceID(ctx context.Context) string {
	traceID, ok := ctx.Value(traceIDKey).(string)
	if !ok || traceID == "" {
		traceID = NewTraceID()
	}
	return traceID
}
