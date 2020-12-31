package logic

import (
	"context"

	"github.com/bjhuang103/go_pay/guarantee_trade/guarantee_trade"
	"github.com/bjhuang103/go_pay/guarantee_trade/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type PrePayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPrePayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrePayLogic {
	return &PrePayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PrePayLogic) PrePay(in *guarantee_trade.PrePayReq) (*guarantee_trade.PrePayResp, error) {
	// todo: add your logic here and delete this line

	return &guarantee_trade.PrePayResp{}, nil
}
