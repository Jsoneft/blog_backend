package service

import (
	"context"
	"ginblog_backend/global"
	"ginblog_backend/internal/dao"
)

type Service struct {
	dao *dao.Dao
	ctx context.Context
}

func New(ctx context.Context) Service {
	service := Service{
		dao: dao.New(global.DBEngine),
		ctx: ctx,
	}
	return service
}
