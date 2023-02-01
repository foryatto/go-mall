package logic

import (
	"context"

	"mall/app/product/rpc/internal/svc"
	"mall/app/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckAndUpdateStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckAndUpdateStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckAndUpdateStockLogic {
	return &CheckAndUpdateStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckAndUpdateStockLogic) CheckAndUpdateStock(in *pb.CheckAndUpdateStockReq) (*pb.CheckAndUpdateStockResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CheckAndUpdateStockResp{}, nil
}
