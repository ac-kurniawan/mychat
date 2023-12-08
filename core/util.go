package core

import "context"

type IUtil interface {
	StartTrace(ctx context.Context, traceName string) (context.Context, any)
	EndTrace(span any)
	TraceError(span any, err error)
	LogInfo(ctx context.Context, message string)
	LogError(ctx context.Context, err error)
}
