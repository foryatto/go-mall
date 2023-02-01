package svc

import (
	"mall/app/user/rpc/internal/config"
	"mall/app/user/rpc/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	//add dependency on user model
	UserModel model.UserModel
	//add dependency on user model
	UserAddressModel    model.UserAddressModel
	UserCollectionModel model.UserCollectionModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:              c,
		UserModel:           model.NewUserModel(sqlConn, c.CacheRedis),
		UserAddressModel:    model.NewUserAddressModel(sqlConn, c.CacheRedis),
		UserCollectionModel: model.NewUserCollectionModel(sqlConn, c.CacheRedis),
	}
}
