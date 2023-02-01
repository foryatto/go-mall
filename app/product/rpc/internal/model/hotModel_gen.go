// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	hotFieldNames          = builder.RawFieldNames(&Hot{})
	hotRows                = strings.Join(hotFieldNames, ",")
	hotRowsExpectAutoSet   = strings.Join(stringx.Remove(hotFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	hotRowsWithPlaceHolder = strings.Join(stringx.Remove(hotFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheMallProductHotIdPrefix = "cache:mallProduct:hot:id:"
)

type (
	hotModel interface {
		Insert(ctx context.Context, data *Hot) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Hot, error)
		Update(ctx context.Context, data *Hot) error
		Delete(ctx context.Context, id int64) error
	}

	defaultHotModel struct {
		sqlc.CachedConn
		table string
	}

	Hot struct {
		Id         int64     `db:"id"`
		ProductId  int64     `db:"product_id"`  // 商品id
		Status     int64     `db:"status"`      // 热门商品状态 0-下线 1-上线
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 更新时间
	}
)

func newHotModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultHotModel {
	return &defaultHotModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`hot`",
	}
}

func (m *defaultHotModel) Delete(ctx context.Context, id int64) error {
	mallProductHotIdKey := fmt.Sprintf("%s%v", cacheMallProductHotIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, mallProductHotIdKey)
	return err
}

func (m *defaultHotModel) FindOne(ctx context.Context, id int64) (*Hot, error) {
	mallProductHotIdKey := fmt.Sprintf("%s%v", cacheMallProductHotIdPrefix, id)
	var resp Hot
	err := m.QueryRowCtx(ctx, &resp, mallProductHotIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", hotRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultHotModel) Insert(ctx context.Context, data *Hot) (sql.Result, error) {
	mallProductHotIdKey := fmt.Sprintf("%s%v", cacheMallProductHotIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, hotRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ProductId, data.Status)
	}, mallProductHotIdKey)
	return ret, err
}

func (m *defaultHotModel) Update(ctx context.Context, data *Hot) error {
	mallProductHotIdKey := fmt.Sprintf("%s%v", cacheMallProductHotIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, hotRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.ProductId, data.Status, data.Id)
	}, mallProductHotIdKey)
	return err
}

func (m *defaultHotModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheMallProductHotIdPrefix, primary)
}

func (m *defaultHotModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", hotRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultHotModel) tableName() string {
	return m.table
}