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

type DelUserCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserCollectionLogic {
	return &DelUserCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除收藏
func (l *DelUserCollectionLogic) DelUserCollection(in *pb.UserCollectionDelReq) (*pb.UserCollectionDelResp, error) {
	_, err := l.svcCtx.UserCollectionModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrap(xerr.NewErrMsg("数据不存在"), "该商品没有被收藏")
		}
		return nil, err
	}
	err = l.svcCtx.UserCollectionModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "DelUserCollection Database Exception : %+v , err: %v", in.Id, err)
	}
	return &user.UserCollectionDelResp{}, nil
}
