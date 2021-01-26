package service

import (
	"context"
	"github.com/bjhuang103/go_pay/common/logs"
	"github.com/bjhuang103/go_pay/pay_core/pay_core"
)

func Pay(ctx context.Context, req *pay_core.PayReq) (*pay_core.PayResp, error) {
	logs.CtxInfo(ctx, "enter pay_service.........")
	return nil, nil
}
