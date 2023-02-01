package logic

import (
	"context"

	"mall/app/user/rpc/internal/model"
	"mall/app/user/rpc/internal/svc"
	"mall/app/user/rpc/pb"
	"mall/app/user/rpc/user"
	"mall/pkg/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type EditUserAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditUserAddressLogic {
	return &EditUserAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 编辑收货地址
func (l *EditUserAddressLogic) EditUserAddress(in *pb.UserAddressEditReq) (*pb.UserAddressEditResp, error) {
	_, err := l.svcCtx.UserAddressModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "数据不存在")
		}
		return nil, err
	}

	//todo 限制,只能有一个默认地址
	// if in.IsDefault == 1 {

	// }

	dbAddress := new(model.UserAddress)
	errCopy := copier.Copy(&dbAddress, in)
	if errCopy != nil {
		return nil, errCopy
	}
	err = l.svcCtx.UserAddressModel.Update(l.ctx, dbAddress)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "AddUserAddress Database Exception : %+v , err: %v", dbAddress, err)
	}
	return &user.UserAddressEditResp{}, nil
}
