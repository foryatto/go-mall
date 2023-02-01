package logic

import (
	"context"
	"fmt"

	"mall/app/product/rpc/internal/model"
	"mall/app/product/rpc/internal/svc"
	"mall/app/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLogic {
	return &ProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取商品信息
func (l *ProductLogic) Product(in *pb.ProductItemReq) (*pb.ProductItem, error) {
	// 防止缓存击穿, 避免缓存过期大量请求涌入服务器
	// 通过singleFlight保证只有一个请求到达服务器
	v, err, _ := l.svcCtx.SingleGroup.Do(fmt.Sprintf("product:%d", in.ProductId), func() (interface{}, error) {
		return l.svcCtx.ProductModel.FindOne(l.ctx, in.ProductId)
	})
	if err != nil {
		return nil, err
	}
	p := v.(*model.Product)
	return &pb.ProductItem{
		ProductId: p.Id,
		Name:      p.Name,
		Stock:     p.Stock,
	}, nil
}
