package logic

import (
	"context"

	"mall/app/user/rpc/internal/model"
	"mall/app/user/rpc/internal/svc"
	"mall/app/user/rpc/pb"
	"mall/app/user/rpc/user"
	"mall/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCollectionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserCollectionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCollectionListLogic {
	return &GetUserCollectionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 收藏列表
func (l *GetUserCollectionListLogic) GetUserCollectionList(in *pb.UserCollectionListReq) (*pb.UserCollectionListResp, error) {
	collectionList, err := l.svcCtx.UserCollectionModel.FindAllByUid(l.ctx, in.Uid)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "Failed  get user's Collection list err : %v , in :%+v", err, in)
	}
	var resp []int64
	for _, collections := range collectionList {
		//_ = copier.Copy(&resp, collections.ProductId)
		resp = append(resp, collections.ProductId)
	}
	return &user.UserCollectionListResp{
		ProductId: resp,
	}, nil
}
