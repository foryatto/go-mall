package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserCollectionModel = (*customUserCollectionModel)(nil)

type (
	// UserCollectionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserCollectionModel.
	UserCollectionModel interface {
		userCollectionModel
		FindAllByUid(ctx context.Context, uid int64) ([]*UserCollection, error)
	}

	customUserCollectionModel struct {
		*defaultUserCollectionModel
	}
)

// NewUserCollectionModel returns a model for the database table.
func NewUserCollectionModel(conn sqlx.SqlConn, c cache.CacheConf) UserCollectionModel {
	return &customUserCollectionModel{
		defaultUserCollectionModel: newUserCollectionModel(conn, c),
	}
}

func (m *defaultUserCollectionModel) FindAllByUid(ctx context.Context, uid int64) ([]*UserCollection, error) {
	var resp []*UserCollection
	query := fmt.Sprintf("select %s from %s where `uid` = ?", userCollectionRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, uid)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
