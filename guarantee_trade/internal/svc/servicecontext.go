package svc

import "github.com/bjhuang103/go_pay/guarantee_trade/internal/config"

type ServiceContext struct {
	c config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c: c,
	}
}
