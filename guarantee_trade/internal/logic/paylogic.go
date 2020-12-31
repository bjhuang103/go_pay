package logic

import (
	"context"

	"github.com/bjhuang103/go_pay/guarantee_trade/guarantee_trade"
	"github.com/bjhuang103/go_pay/guarantee_trade/internal/svc"

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

func (l *PayLogic) Pay(in *guarantee_trade.PayReq) (*guarantee_trade.PayResp, error) {
	// todo: add your logic here and delete this line

	return &guarantee_trade.PayResp{}, nil
}
