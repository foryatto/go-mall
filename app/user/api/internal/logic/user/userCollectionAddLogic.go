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

type UserCollectionAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCollectionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCollectionAddLogic {
	return &UserCollectionAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCollectionAddLogic) UserCollectionAdd(req *types.UserCollectionAddReq) (resp *types.UserCollectionAddRes, err error) {
	//get uid from token
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()

	var addRpcReq user.UserCollectionAddReq
	addRpcReq.Uid = uid
	errCopy := copier.Copy(&addRpcReq, req)
	if err != nil {
		return nil, errCopy
	}

	_, err = l.svcCtx.UserRPC.AddUserCollection(l.ctx, &addRpcReq)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
