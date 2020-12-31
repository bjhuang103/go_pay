package logic

import (
	"context"
	"github.com/bjhuang103/go_pay/pay_core/service"

	"github.com/bjhuang103/go_pay/pay_core/internal/svc"
	"github.com/bjhuang103/go_pay/pay_core/pay_core"

	"github.com/tal-tech/go-zero/core/logx"
)

type PayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayLogic {
	return &PayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PayLogic) Pay(in *pay_core.PayReq) (*pay_core.PayResp, error) {
	l.Info("enter rpc pay...")

	return service.Pay(l.ctx, in)
}
