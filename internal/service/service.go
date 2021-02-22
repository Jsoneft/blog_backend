package service

import (
	"context"
	"ginblog_backend/global"
	"ginblog_backend/internal/dao"
	otgorm "github.com/eddycjy/opentracing-gorm"
)

type Service struct {
	dao *dao.Dao
	ctx context.Context
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
	return svc
}
