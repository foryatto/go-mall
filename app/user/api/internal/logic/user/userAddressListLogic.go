package user

import (
	"context"
	"encoding/json"

	"mall/app/user/api/internal/svc"
	"mall/app/user/api/internal/types"
	"mall/app/user/rpc/user"
	"mall/pkg/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddressListLogic {
	return &UserAddressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddressListLogic) UserAddressList(req *types.UserAddressListReq) (resp *types.UserAddressListRes, err error) {
	var addressListReq user.UserAddressListReq
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("Error! get uid from token"), "Failed toget uid from token err : %v ,req:%+v", err, req)
	}
	addressListReq.Uid = uid
	rpcRes, err := l.svcCtx.UserRPC.GetUserAddressList(l.ctx, &addressListReq)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("Error! Function UserReceiveAddressList"), "Failed to get user addrerss  list err : %v ,req:%+v", err, req)
	}
	var addressList []types.UserAddress
	for _, rpcAddress := range rpcRes.List {
		var addressVo types.UserAddress
		_ = copier.Copy(&addressVo, rpcAddress)
		addressList = append(addressList, addressVo)
	}
	return &types.UserAddressListRes{List: addressList}, nil
}
