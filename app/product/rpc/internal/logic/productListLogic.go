package logic

import (
	"context"

	"mall/app/product/rpc/internal/svc"
	"mall/app/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListLogic {
	return &ProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductListLogic) ProductList(in *pb.ProductListReq) (*pb.ProductListResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ProductListResp{}, nil
}
