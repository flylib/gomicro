package handler

import (
	"context"
	"crypto/md5"
	"fmt"
	"hello/proto"
)

// 业务实现方法的容器
type MD5Handler struct{}

// 为server定义 DoMD5 方法 内部处理请求并返回结果
// 参数 (context.Context[固定], *proto.Req[相应接口定义的请求参数])
// 返回 (*proto.Res[相应接口定义的返回参数，必须用指针], error)
func (s *MD5Handler) DoMD5(ctx context.Context, in *proto.Req) (*proto.Res, error) {
	fmt.Println("[DoMD5]recv ", in.JsonStr)
	return &proto.Res{BackJson: "MD5 :" + fmt.Sprintf("%x", md5.Sum([]byte(in.JsonStr)))}, nil
}
