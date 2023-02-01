package user

import (
	"context"

	"mall/app/user/api/internal/svc"
	"mall/app/user/api/internal/types"
	"mall/app/user/rpc/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type EditAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditAddressLogic {
	return &EditAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditAddressLogic) EditAddress(req *types.UserAddressEditReq) (resp *types.UserAddressEditRes, err error) {
	var editRpcReq user.UserAddressEditReq
	errCopy := copier.Copy(&editRpcReq, req)
	if errCopy != nil {
		return nil, errCopy
	}
	_, err = l.svcCtx.UserRPC.EditUserAddress(l.ctx, &editRpcReq)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
