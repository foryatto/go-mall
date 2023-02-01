package logic

import (
	"context"
	"strconv"
	"strings"

	"mall/app/product/rpc/internal/model"
	"mall/app/product/rpc/internal/svc"
	"mall/app/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type ProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductsLogic {
	return &ProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductsLogic) Products(in *pb.ProductReq) (*pb.ProductResp, error) {
	products := make(map[int64]*pb.ProductItem)
	pdis := strings.Split(in.ProductIds, ",")
	// 用gozero的mapReduce做并发控制,如果有一个出错能立即返回
	// 生产, 处理, 聚合
	ps, err := mr.MapReduce(func(source chan<- interface{}) {
		for _, pid := range pdis {
			source <- pid
		}
	}, func(item interface{}, writer mr.Writer, cancel func(error)) {
		pidStr := item.(string)
		pid, err := strconv.ParseInt(pidStr, 10, 64)
		if err != nil {
			return
		}
		p, err := l.svcCtx.ProductModel.FindOne(l.ctx, pid)
		if err != nil {
			return
		}
		writer.Write(p)
	}, func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
		var r []*model.Product
		for p := range pipe {
			r = append(r, p.(*model.Product))
		}
		writer.Write(r)
	})
	if err != nil {
		return nil, err
	}
	for _, p := range ps.([]*model.Product) {
		products[p.Id] = &pb.ProductItem{
			ProductId: p.Id,
			Name:      p.Name,
			Price:     p.Price,
		}
	}
	return &pb.ProductResp{Products: products}, nil
}
