package utils

import (
	"context"
)

const (
	HeaderTraceId = "x-trace-id"
)

func GetTrxId(ctx context.Context) string {
	// Get trxId from the context
	trxId, ok := ctx.Value(HeaderTraceId).(string)

	if ok {
		return trxId
	}

	// If trxId is not found or is of the wrong type, return an empty string or handle as needed
	return ""
}

func SetCustomMetaDataTransactionId(ctx context.Context, trxid string) context.Context {
	if ctx == nil {
		return context.Background()
	}
	return context.WithValue(ctx, HeaderTraceId, trxid)
}
