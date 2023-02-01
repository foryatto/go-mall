package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HotModel = (*customHotModel)(nil)

type (
	// HotModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHotModel.
	HotModel interface {
		hotModel
	}

	customHotModel struct {
		*defaultHotModel
	}
)

// NewHotModel returns a model for the database table.
func NewHotModel(conn sqlx.SqlConn, c cache.CacheConf) HotModel {
	return &customHotModel{
		defaultHotModel: newHotModel(conn, c),
	}
}
