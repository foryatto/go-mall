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
	categoryFieldNames          = builder.RawFieldNames(&Category{})
	categoryRows                = strings.Join(categoryFieldNames, ",")
	categoryRowsExpectAutoSet   = strings.Join(stringx.Remove(categoryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	categoryRowsWithPlaceHolder = strings.Join(stringx.Remove(categoryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheMallProductCategoryIdPrefix = "cache:mallProduct:category:id:"
)

type (
	categoryModel interface {
		Insert(ctx context.Context, data *Category) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Category, error)
		Update(ctx context.Context, data *Category) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCategoryModel struct {
		sqlc.CachedConn
		table string
	}

	Category struct {
		Id         int64     `db:"id"`          // 分类id
		Parentid   int64     `db:"parentid"`    // 父类别id当id=0时说明是根节点,一级类别
		Name       string    `db:"name"`        // 类别名称
		Status     int64     `db:"status"`      // 类别状态1-正常,2-已废弃
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 更新时间
	}
)

func newCategoryModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultCategoryModel {
	return &defaultCategoryModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`category`",
	}
}

func (m *defaultCategoryModel) Delete(ctx context.Context, id int64) error {
	mallProductCategoryIdKey := fmt.Sprintf("%s%v", cacheMallProductCategoryIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, mallProductCategoryIdKey)
	return err
}

func (m *defaultCategoryModel) FindOne(ctx context.Context, id int64) (*Category, error) {
	mallProductCategoryIdKey := fmt.Sprintf("%s%v", cacheMallProductCategoryIdPrefix, id)
	var resp Category
	err := m.QueryRowCtx(ctx, &resp, mallProductCategoryIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", categoryRows, m.table)
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

func (m *defaultCategoryModel) Insert(ctx context.Context, data *Category) (sql.Result, error) {
	mallProductCategoryIdKey := fmt.Sprintf("%s%v", cacheMallProductCategoryIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, categoryRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Parentid, data.Name, data.Status)
	}, mallProductCategoryIdKey)
	return ret, err
}

func (m *defaultCategoryModel) Update(ctx context.Context, data *Category) error {
	mallProductCategoryIdKey := fmt.Sprintf("%s%v", cacheMallProductCategoryIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, categoryRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Parentid, data.Name, data.Status, data.Id)
	}, mallProductCategoryIdKey)
	return err
}

func (m *defaultCategoryModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheMallProductCategoryIdPrefix, primary)
}

func (m *defaultCategoryModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", categoryRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultCategoryModel) tableName() string {
	return m.table
}
