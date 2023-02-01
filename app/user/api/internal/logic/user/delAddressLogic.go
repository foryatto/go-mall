package user

import (
	"context"

	"mall/app/user/api/internal/svc"
	"mall/app/user/api/internal/types"
	"mall/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAddressLogic {
	return &DelAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelAddressLogic) DelAddress(req *types.UserAddressDelReq) (resp *types.UserAddressDelRes, err error) {
	_, err = l.svcCtx.UserRPC.DelUserAddress(l.ctx, &user.UserAddressDelReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
