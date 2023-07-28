package grpc

import (
	"context"
	"google.golang.org/grpc/stats"
	"log"
)

type StatsHandler struct {
}

//TagConn可以将一些信息附加到给定的上下文。
func (h *StatsHandler) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	log.Println("tagConn...", info.RemoteAddr)
	return ctx
}

// 会在连接开始和结束时被调用，分别会输入不同的状态.
func (h *StatsHandler) HandleConn(ctx context.Context, s stats.ConnStats) {
	// 开始和结束状态
	switch s.(type) {
	case *stats.ConnBegin:
		log.Printf("begin conn")
	case *stats.ConnEnd:
		log.Printf("end conn")
	default:
		log.Println("handleConn...")
	}
}

// TagRPC可以将一些信息附加到给定的上下文
func (h *StatsHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	log.Println("tagrpc...@" + info.FullMethodName)
	return ctx
}

// 处理RPC统计信息
func (h *StatsHandler) HandleRPC(ctx context.Context, s stats.RPCStats) {
	switch s.(type) {
	case *stats.Begin:
		log.Println("handlerRPC begin...")
	case *stats.End:
		log.Println("handlerRPC End...")
	case *stats.InHeader:
		log.Println("handlerRPC InHeader...")
	case *stats.InPayload:
		log.Println("handlerRPC InPayload...")
	case *stats.InTrailer:
		log.Println("handlerRPC InTrailer...")
	case *stats.OutHeader:
		log.Println("handlerRPC OutHeader...")
	case *stats.OutPayload:
		log.Println("handlerRPC OutPayload...")
	default:
		log.Println("handleRPC...")
	}
}
