package controller

import (
	"context"
	"github.com/lie-flat-planet/srv-kubelie/config"
	"github.com/lie-flat-planet/srv-kubelie/internal/model"
	"github.com/lie-flat-planet/srv-kubelie/internal/service"
)

func FactoryDemo() *Demo {
	return &Demo{
		service: service.NewDemo(
			service.WithMysql(config.Config.Mysql.GetDB()),
			service.WithRedis(config.Config.Redis.GetClient()),
		),
	}
}

type Demo struct {
	service *service.Demo
}

func (demo *Demo) FetchFirst(ctx context.Context) (*model.User, error) {
	return demo.service.FetchFirst(ctx)
}