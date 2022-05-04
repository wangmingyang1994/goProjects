package service

import (
	"context"
	"goProjects/blog_service/global"
	"goProjects/blog_service/internal/dao"
)

type Service struct{
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context)Service{
	svc :=Service{
		ctx:ctx,
	}
	svc.dao=dao.New(global.DBEngine)
	return svc
}

