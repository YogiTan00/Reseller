package utils

import (
	"context"
	"google.golang.org/grpc/metadata"
)

const (
	HeaderTraceId = "x-trace-id"
)

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
