package logic

import (
	"context"

	"github.com/bjhuang103/go_pay/pay_core/internal/svc"
	"github.com/bjhuang103/go_pay/pay_core/pay_core"

	"github.com/tal-tech/go-zero/core/logx"
)

type RefundLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefundLogic {
	return &RefundLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefundLogic) Refund(in *pay_core.RefundReq) (*pay_core.RefundResp, error) {
	// todo: add your logic here and delete this line

	return &pay_core.RefundResp{}, nil
}
