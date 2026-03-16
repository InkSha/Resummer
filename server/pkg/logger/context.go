package logger

import (
	"context"
	"math/rand/v2"
	"strconv"
	"time"
)

func WithTraceID(ctx context.Context) context.Context {
	traceID := generateTraceID()
	return context.WithValue(ctx, "trace_id", traceID)
}

func generateTraceID() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10) + strconv.Itoa(rand.IntN(1000))
}
