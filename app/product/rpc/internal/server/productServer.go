// Code generated by goctl. DO NOT EDIT.
// Source: product.proto

package server

import (
	"context"

	"mall/app/product/rpc/internal/logic"
	"mall/app/product/rpc/internal/svc"
	"mall/app/product/rpc/pb"
)

type ProductServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedProductServer
}

func NewProductServer(svcCtx *svc.ServiceContext) *ProductServer {
	return &ProductServer{
		svcCtx: svcCtx,
	}
}

func (s *ProductServer) Product(ctx context.Context, in *pb.ProductItemReq) (*pb.ProductItem, error) {
	l := logic.NewProductLogic(ctx, s.svcCtx)
	return l.Product(in)
}

func (s *ProductServer) Products(ctx context.Context, in *pb.ProductReq) (*pb.ProductResp, error) {
	l := logic.NewProductsLogic(ctx, s.svcCtx)
	return l.Products(in)
}

func (s *ProductServer) ProductList(ctx context.Context, in *pb.ProductListReq) (*pb.ProductListResp, error) {
	l := logic.NewProductListLogic(ctx, s.svcCtx)
	return l.ProductList(in)
}

func (s *ProductServer) HotProducts(ctx context.Context, in *pb.HotProductsReq) (*pb.HotProductsResp, error) {
	l := logic.NewHotProductsLogic(ctx, s.svcCtx)
	return l.HotProducts(in)
}

func (s *ProductServer) UpdateProductStock(ctx context.Context, in *pb.UpdateProductStockReq) (*pb.UpdateProductStockResp, error) {
	l := logic.NewUpdateProductStockLogic(ctx, s.svcCtx)
	return l.UpdateProductStock(in)
}

func (s *ProductServer) CheckAndUpdateStock(ctx context.Context, in *pb.CheckAndUpdateStockReq) (*pb.CheckAndUpdateStockResp, error) {
	l := logic.NewCheckAndUpdateStockLogic(ctx, s.svcCtx)
	return l.CheckAndUpdateStock(in)
}

func (s *ProductServer) CheckProductStock(ctx context.Context, in *pb.UpdateProductStockReq) (*pb.UpdateProductStockResp, error) {
	l := logic.NewCheckProductStockLogic(ctx, s.svcCtx)
	return l.CheckProductStock(in)
}

func (s *ProductServer) RollbackProductStock(ctx context.Context, in *pb.UpdateProductStockReq) (*pb.UpdateProductStockResp, error) {
	l := logic.NewRollbackProductStockLogic(ctx, s.svcCtx)
	return l.RollbackProductStock(in)
}

func (s *ProductServer) DecrStock(ctx context.Context, in *pb.DecrStockReq) (*pb.DecrStockResp, error) {
	l := logic.NewDecrStockLogic(ctx, s.svcCtx)
	return l.DecrStock(in)
}

func (s *ProductServer) DecrStockRevert(ctx context.Context, in *pb.DecrStockReq) (*pb.DecrStockResp, error) {
	l := logic.NewDecrStockRevertLogic(ctx, s.svcCtx)
	return l.DecrStockRevert(in)
}
