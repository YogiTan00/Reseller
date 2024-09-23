package utils

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
)

const (
	HeaderTraceId = "x-trace-id"
)

func GetTrxId(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	txID, ok := ctx.Value(HeaderTraceId).(string)
	if !ok {
		fmt.Println("Transaction ID not found in context")
		return ""
	}
	return txID
}

func GetMDWithTrxid(ctx context.Context, trxid string) metadata.MD {
	if ctx == nil {
		return metadata.MD{}
	}
	mds, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if v := mds.Get(HeaderTraceId); len(v) > 0 {
			if v[0] == "" {
				mds.Set(HeaderTraceId, trxid)
				return mds
			}
			return mds
		}
		mds.Append(HeaderTraceId, trxid)
		return mds
	}
	return metadata.Pairs(HeaderTraceId, trxid)
}

func SetCustomMetaDataTransactionId(ctx context.Context, trxid string) context.Context {
	if ctx == nil {
		return context.Background()
	}
	mds := GetMDWithTrxid(ctx, trxid)
	return metadata.NewIncomingContext(ctx, mds)
}
