package user

import (
	"context"
	"encoding/json"

	"mall/app/user/api/internal/svc"
	"mall/app/user/api/internal/types"
	"mall/app/user/rpc/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAddressLogic {
	return &AddAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddAddressLogic) AddAddress(req *types.UserAddressAddReq) (resp *types.UserAddressAddRes, err error) {
	//get uid from jwt token
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	var addRpcReq user.UserAddressAddReq
	addRpcReq.Uid = uid
	errCopy := copier.Copy(&addRpcReq, req)
	if errCopy != nil {
		return nil, errCopy
	}
	_, err = l.svcCtx.UserRPC.AddUserAddress(l.ctx, &addRpcReq)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
