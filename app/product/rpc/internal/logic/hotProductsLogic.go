package logic

import (
	"context"

	"mall/app/product/rpc/internal/svc"
	"mall/app/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type HotProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHotProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HotProductsLogic {
	return &HotProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HotProductsLogic) HotProducts(in *pb.HotProductsReq) (*pb.HotProductsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.HotProductsResp{}, nil
}
